package router

import (
	"net/http"

	"sky-take-out-go/controller/admin"
	"sky-take-out-go/controller/rider"
	"sky-take-out-go/controller/user"
	"sky-take-out-go/middleware"
	"sky-take-out-go/websocket"

	"github.com/gin-gonic/gin"
	gorillaWs "github.com/gorilla/websocket"
)

var upgrader = gorillaWs.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}

func Setup() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.CORS())

	// admin routes (JWT auth)
	a := r.Group("/admin")
	a.POST("/employee/login", admin.Login)
	a.Use(middleware.AdminAuth())
	{
		a.POST("/employee/logout", admin.Logout)
		a.POST("/employee", admin.AddEmployee)
		a.GET("/employee/page", admin.EmployeePage)
		a.GET("/employee/:id", admin.GetEmployee)
		a.PUT("/employee", admin.UpdateEmployee)
		a.POST("/employee/status/:status", admin.SetEmployeeStatus)
		a.PUT("/employee/editPassword", admin.EditPassword)
	}
	{
		a.POST("/category", admin.AddCategory)
		a.GET("/category/page", admin.CategoryPage)
		a.DELETE("/category", admin.DeleteCategory)
		a.PUT("/category", admin.UpdateCategory)
		a.POST("/category/status/:status", admin.SetCategoryStatus)
		a.GET("/category/list", admin.CategoryList)
	}
	{
		a.POST("/dish", admin.AddDish)
		a.GET("/dish/page", admin.DishPage)
		a.DELETE("/dish", admin.DeleteDish)
		a.GET("/dish/:id", admin.GetDish)
		a.PUT("/dish", admin.UpdateDish)
		a.POST("/dish/status/:status", admin.SetDishStatus)
		a.GET("/dish/list", admin.DishList)
	}
	{
		a.POST("/setmeal", admin.AddSetmeal)
		a.GET("/setmeal/page", admin.SetmealPage)
		a.DELETE("/setmeal", admin.DeleteSetmeal)
		a.GET("/setmeal/:id", admin.GetSetmeal)
		a.PUT("/setmeal", admin.UpdateSetmeal)
		a.POST("/setmeal/status/:status", admin.SetSetmealStatus)
		a.GET("/setmeal/list", admin.SetmealList)
		a.GET("/setmeal/dish/:id", admin.GetSetmealDishes)
	}
	{
		a.GET("/order/conditionSearch", admin.OrderConditionSearch)
		a.GET("/order/statistics", admin.OrderStatistics)
		a.GET("/order/details/:id", admin.OrderDetails)
		a.PUT("/order/confirm", admin.ConfirmOrder)
		a.PUT("/order/rejection", admin.RejectOrder)
		a.PUT("/order/cancel", admin.CancelOrder)
		a.PUT("/order/delivery/:id", admin.DeliveryOrder)
		a.PUT("/order/complete/:id", admin.CompleteOrder)
	}
	{
		a.GET("/shop/status", admin.GetShopStatus)
		a.PUT("/shop/status/:status", admin.SetShopStatus)
	}
	{
		a.POST("/common/upload", admin.Upload)
	}
	{
		a.GET("/report/turnoverStatistics", admin.TurnoverStatistics)
		a.GET("/report/userStatistics", admin.UserStatistics)
		a.GET("/report/ordersStatistics", admin.OrdersStatistics)
		a.GET("/report/top10", admin.Top10)
		a.GET("/report/export", admin.Export)
	}
	{
		a.GET("/workspace/businessData", admin.BusinessData)
		a.GET("/workspace/overviewOrders", admin.OverviewOrders)
		a.GET("/workspace/overviewDishes", admin.OverviewDishes)
		a.GET("/workspace/overviewSetmeals", admin.OverviewSetmeals)
	}

	// user routes (JWT auth)
	u := r.Group("/user")
	u.POST("/user/login", user.Login)
	u.Use(middleware.UserAuth())
	{
		u.POST("/user/logout", user.Logout)
	}
	{
		u.GET("/category/list", user.CategoryList)
	}
	{
		u.GET("/dish/list", user.DishList)
	}
	{
		u.GET("/setmeal/list", user.SetmealList)
		u.GET("/setmeal/dish/:id", user.GetSetmealDishes)
	}
	{
		u.GET("/shop/status", user.GetShopStatus)
	}
	{
		u.POST("/shoppingCart/add", user.AddToCart)
		u.GET("/shoppingCart/list", user.CartList)
		u.DELETE("/shoppingCart/clean", user.CleanCart)
		u.POST("/shoppingCart/sub", user.SubFromCart)
	}
	{
		u.POST("/order/submit", user.SubmitOrder)
		u.PUT("/order/payment", user.Payment)
		u.GET("/order/historyOrders", user.HistoryOrders)
		u.GET("/order/orderDetail/:id", user.OrderDetail)
		u.PUT("/order/cancel/:id", user.CancelOrder)
		u.POST("/order/repetition/:id", user.RepetitionOrder)
		u.GET("/order/reminder/:id", user.Reminder)
	}
	{
		u.POST("/addressBook", user.AddAddress)
		u.PUT("/addressBook", user.UpdateAddress)
		u.GET("/addressBook/list", user.AddressList)
		u.GET("/addressBook/:id", user.GetAddress)
		u.DELETE("/addressBook", user.DeleteAddress)
		u.PUT("/addressBook/default", user.SetDefaultAddress)
		u.GET("/addressBook/default", user.GetDefaultAddress)
	}
	// static
	r.Static("/uploads", "./uploads")

	// rider API routes
	ri := r.Group("/rider")
	ri.POST("/login", rider.Login)
	ri.Use(middleware.AdminAuth())
	{
		ri.GET("/orders/pending", rider.PendingOrders)
		ri.GET("/orders/grab", rider.GrabOrder)
		ri.GET("/orders/my", rider.MyOrders)
		ri.GET("/orders/complete", rider.CompleteOrder)
	}

	// rider static pages
	r.GET("/rider/login.html", func(c *gin.Context) { c.File("../rider/login.html") })
	r.GET("/rider/orders.html", func(c *gin.Context) { c.File("../rider/orders.html") })

	// websocket
	r.GET("/ws/:sid", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		sid := c.Param("sid")
		websocket.DefaultHub.HandleConn(conn, sid)
	})

	return r
}
