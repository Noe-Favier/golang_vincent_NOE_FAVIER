package routers

import (
	middlewares "blog_noe/middlewares"

	privateHandlers "blog_noe/handlers/private"
	publicHandlers "blog_noe/handlers/public"

	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
)

type Routes struct {
	router *gin.Engine
}

func New() *Routes {
	return &Routes{
		router: gin.New(),
	}
}

func (r *Routes) SetupRouter() *gin.Engine {
	// Middlewares de base
	r.router.Use(gin.Recovery())
	r.router.Use(gin.Logger())

	//	Security headers
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect:           false,
		STSSeconds:            315360000,
		STSIncludeSubdomains:  true,
		FrameDeny:             true,
		ContentTypeNosniff:    true,
		BrowserXssFilter:      true,
		ContentSecurityPolicy: "default-src 'self'; style-src 'self' 'unsafe-inline'; script-src 'self' 'unsafe-inline'",
		ReferrerPolicy:        "strict-origin-when-cross-origin",
	})

	r.router.Use(func(c *gin.Context) {
		err := secureMiddleware.Process(c.Writer, c.Request)
		if err != nil {
			c.Abort()
			return
		}
		c.Next()
	})

	// Routes setup
	// public
	r.setupPublicRoutes()

	// private
	r.setupProtectedRoutes()

	return r.router
}

func (r *Routes) setupPublicRoutes() {
	public := r.router.Group("/api/public")
	{
		public.GET("/health", publicHandlers.HealthHandler)
		public.GET("/home", publicHandlers.HomeHandler)
		public.POST("/login", publicHandlers.Login)
	}
}

func (r *Routes) setupProtectedRoutes() {
	protected := r.router.Group("/api/")
	protected.Use(middlewares.AuthMiddleware())
	{
		protected.GET("/health", privateHandlers.HealthHandler)

		//posts
		protected.GET("/posts", privateHandlers.UserPostsByMailHandler)
		protected.POST("/posts", privateHandlers.UploadPostHandler)
	}
}
