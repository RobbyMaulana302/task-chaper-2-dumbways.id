package middleware

// func UploadFile(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		file, err := c.FormFile("input-image")
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 		}

// 		src, err := file.Open()
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 		}
// 		defer src.Close()

// 		tempFile, err := ioutil.TempFile("uploads", "image-*.png")
// 		if err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 		}
// 		defer tempFile.Close()

// 		if _, err := io.Copy(tempFile, src); err != nil {
// 			return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
// 		}

// 		data := tempFile.Name()
// 		filename := data[8:]

// 		c.Set("dataFile", filename)

// 		return next(c)
// 	}
// }