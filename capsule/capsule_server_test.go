package capsule

type serverTestCase struct {
	name                    string
	input                   string
	expectedStatusCode      int
	expectedResponseMessage string
}

/*
func TestCapsuleServer(t *testing.T) {
	_ := []serverTestCase{
		{
			name:                    "Valid request",
			input:                   "{ \"type\": \"email\", \"from\": \"jon@email.com\", \"to\": \"jane@email.com\", \"deliveryDate\": \"9/30/2123\"}",
			expectedStatusCode:      200,
			expectedResponseMessage: "",
		},
		{
			name:                    "Invalid type request",
			input:                   "{ \"type\": \"invalid\", \"from\": \"jon@email.com\", \"to\": \"jane@email.com\", \"deliveryDate\": \"9/30/2123\"}",
			expectedStatusCode:      400,
			expectedResponseMessage: "Invalid capsule type",
		},
		{
			name:                    "Invalid email address request",
			input:                   "{ \"type\": \"invalid\", \"from\": \"jon@emailcom\", \"to\": \"jane@email.com\", \"deliveryDate\": \"9/30/2123\"}",
			expectedStatusCode:      400,
			expectedResponseMessage: "Invalid To or From fields",
		},
		{
			name:                    "Invalid email address request",
			input:                   "{ \"type\": \"invalid\", \"from\": \"jon@emailcom\", \"to\": \"jane@email.com\", \"deliveryDate\": \"9/30/2123\"}",
			expectedStatusCode:      400,
			expectedResponseMessage: "Invalid To or From fields",
		},
	}

		requestWaitGroup := new(sync.WaitGroup)
		baseLogger, err := zap.NewDevelopment()
		if err != nil {
			log.Fatalf("Unable to intialize zap logger: %s\n", err.Error())
		}
		logger := baseLogger.Sugar()
		middleware := capsule.Pipeline(requestWaitGroup, logger)
		service := capsule.Service(logger)
		controller := capsule.Controller(service, logger)
		handler := middleware.ValidateRequest(middleware.ParseCapsuleFromRequest(controller))
		server := httptest.NewServer(handler)
		for _, testCase := range testCases {
			t.Run(testCase.name, func(t *testing.T) {
				body := generateRequestBody(testCase.input)
				request := httptest.NewRequest(http.MethodPost, "/", body)
				writer := httptest.NewRecorder()

			})
		}

}
*/
