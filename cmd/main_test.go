package main

// func Test(t *testing.T) {
// 	req, err := http.NewRequest("Get", "/ejemplo", nil)

// 	if err != nil {
// 		t.Fatalf("could not %v", err)
// 	}

// 	t.Logf("res %v", req)
// }

// func TestNewHTTPServer(t *testing.T) {
// 	type args struct {
// 		lc     fx.Lifecycle
// 		app    *fiber.App
// 		conn   *mongo.Connection
// 		cancel context.CancelFunc
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 	}{}

// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			NewHTTPServer(tt.args.lc, tt.args.app, tt.args.conn, tt.args.cancel)
// 		})
// 	}
// }

// func TestEj(t *testing.T) {
// 	tests := []struct {
// 		description  string // description of the test case
// 		route        string // route path to test
// 		expectedCode int    // expected HTTP status code
// 	}{
// 		// First test case
// 		{
// 			description:  "get HTTP status 200",
// 			route:        "/hello",
// 			expectedCode: 200,
// 		},
// 		// Second test case
// 		{
// 			description:  "get HTTP status 404, when route is not exists",
// 			route:        "/not-found",
// 			expectedCode: 404,
// 		},
// 	}
// 	app := fiber.New()

// 	// Create route with GET method for test
// 	app.Get("/hello", func(c *fiber.Ctx) error {
// 		// Return simple string as response
// 		return c.SendString("Hello, World!")
// 	})

// 	for _, test := range tests {
// 		req := httptest.NewRequest("GET", test.route, nil)
// 		res, _ := app.Test(req, 1)

// 		//assert.Equalf()
// 		assert.Equalf(t, test.expectedCode, res.StatusCode, test.description)
// 	}
// }
