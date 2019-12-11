package router

import (
	"db-forum/api"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
)


func postsRoute(ctx *fasthttp.RequestCtx) {
	options := ctx.UserValue("options").(string)
	if options == "/create" {
		api.CreateForum(ctx)
		return
	}

	options = options[1 : len(options) - 7]
	api.CreateThread(ctx, options)
}

func CreateRouter() *fasthttprouter.Router {
	rt := fasthttprouter.New()

	rt.POST("/api/user/:nickname/create", api.CreateUser)
	rt.POST("/api/user/:nickname/profile", api.UpdateUser)
	rt.GET("/api/user/:nickname/profile", api.GetUser)

	rt.POST("/api/forum/*options", postsRoute)
	rt.GET("/api/forum/:slug/threads", api.GetForumThreads)
	rt.GET("/api/forum/:slug/users", api.GetForumUsers)
	rt.GET("/api/forum/:slug/details", api.GetForum)

	rt.POST("/api/thread/:slug/details", api.UpdateThread)
	rt.POST("/api/thread/:slug/create", api.CreatePost)
	rt.GET("/api/thread/:slug/details", api.GetThread)
	rt.GET("/api/thread/:slug", api.GetThread)
	rt.POST("/api/thread/:slug/vote", api.VoteThread)

	rt.POST("/api/post/:slug/details", api.UpdatePost)
	rt.GET("/api/thread/:slug/posts", api.GetPost)
	rt.GET("/api/post/:slug/details", api.GetPostDetails)

	rt.POST("/api/service/clear", api.ClearService)
	rt.GET("/api/service/status", api.GetServiceStatus)
	return rt
}
