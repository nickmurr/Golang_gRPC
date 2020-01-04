package crudserver

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go_grpc_server/blog/blogpb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gopkg.in/mgo.v2/bson"
	"log"
)

var Collection *mongo.Collection

type CrudServer struct{}

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id,omitempty" json:"_id,omitempty"`
	AuthorID string             `bson:"author_id,omitempty" json:"author_id,omitempty"`
	Content  string             `bson:"content,omitempty" json:"content,omitempty"`
	Title    string             `bson:"title,omitempty" json:"title,omitempty"`
}

func (CrudServer) CreateBlog(ctx context.Context, req *blogpb.CreateBlogRequest) (*blogpb.CreateBlogResponse, error) {
	blog := req.GetBlog()

	data := BlogItem{
		AuthorID: blog.GetAuthor(),
		Title:    blog.GetTitle(),
		Content:  blog.GetContent(),
	}

	result, err := Collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Cannot convert to OID"),
		)
	}

	return &blogpb.CreateBlogResponse{
		Blog: &blogpb.Blog{
			Id:      oid.Hex(),
			Author:  blog.GetAuthor(),
			Title:   blog.GetTitle(),
			Content: blog.GetContent(),
		},
	}, nil
}

func (CrudServer) ReadBlog(ctx context.Context, req *blogpb.ReadBlogRequest) (*blogpb.ReadBlogResponse, error) {
	blogId := req.GetBlogId()

	oid, err := primitive.ObjectIDFromHex(blogId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Cannot parse ID"))
	}

	// 	Create an empty struct
	data := &BlogItem{}

	filter := bson.M{"_id": oid,}

	blogItem := Collection.FindOne(context.Background(), filter)

	if err := blogItem.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Cannot find blog with specified ID: %v", err),
		)
	}

	return &blogpb.ReadBlogResponse{
		Blog: dataToBlogPb(data),
	}, nil
}

func (CrudServer) ReadAllBlog(ctx context.Context, req *blogpb.ReadAllBlogRequest) (*blogpb.ReadAllBlogResponse, error) {
	blogPage := req.GetPage()
	blogSearch := req.GetSearch()

	if blogPage < 0 || blogPage == 0 {
		blogPage = 0
	}

	//

	skip := int64((blogPage - 1) * 10)
	limit := int64(10)

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)

	countOptions := options.CountOptions{
		Skip:  &skip,
		Limit: &limit,
	}

	var data []*blogpb.Blog

	filter := bson.M{"title": primitive.Regex{Pattern: blogSearch}}

	blogItem, err := Collection.Find(context.Background(), filter, findOptions)
	countDocuments, err := Collection.CountDocuments(context.Background(), filter, &countOptions)

	if err != nil {
		log.Fatalf("Error while reading blog %v", err)
	}
	defer blogItem.Close(context.TODO())

	for blogItem.Next(context.TODO()) {
		var blog BlogItem
		errMsg := blogItem.Decode(&blog)
		fmt.Println()
		if errMsg != nil {
			log.Fatal(errMsg)
		}
		data = append(data, dataToBlogPb(&blog))
	}

	log.Println(data, countDocuments)

	return &blogpb.ReadAllBlogResponse{Blog: data}, nil
}

func dataToBlogPb(data *BlogItem) *blogpb.Blog {
	return &blogpb.Blog{
		Id:      data.ID.Hex(),
		Author:  data.AuthorID,
		Content: data.Content,
		Title:   data.Title,
	}
}
