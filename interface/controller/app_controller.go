package controller

type AppController struct {
	Centre interface{ CentreController }
}
