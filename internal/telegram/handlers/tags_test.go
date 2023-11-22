package handlers

//type MockPinterestAPI struct {
//	mock.Mock
//}
//
//func (m *MockPinterestAPI) GetPinsBySearch(query string) (*[]Pin, error) {
//	args := m.Called(query)
//	return args.Get(0).(*[]Pin), args.Error(1)
//}
//
//
//func TestTagsCommand_GetImageList(t *testing.T) {
//	tags := &TagsCommand{
//		UserImageStore: &UserImageStore{
//			ImageLists: map[int32][]string{},
//		},
//	}
//
//	update := telego.Update{
//		Message: &telego.Message{
//			Chat: telego.Chat{
//				ID: 123,
//			},
//			Text: "your_test_tag",
//		},
//	}
//
//	imageList, err := tags.GetImageList(update)
//
//	assert.NoError(t, err)
//	assert.NotEmpty(t, imageList)
//}
//
//func TestTagsCommand_handleUserMessage(t *testing.T) {
//	tags := &TagsCommand{
//		UserImageStore: &UserImageStore{
//			ImageLists: map[int32][]string{},
//		},
//	}
//
//	update := telego.Update{
//		Message: &telego.Message{
//			Text: "user_input",
//		},
//	}
//
//	user, err := tags.handleUserMessage(update)
//
//	assert.NoError(t, err)
//	assert.Equal(t, "user_input", user)
//}
//
//func TestTagsCommand_handleImage(t *testing.T) {
//	tags := &TagsCommand{
//	}
//
//	mockPinterestAPI := new(MockPinterestAPI)
//	tags.Services.PinterestAPI = mockPinterestAPI
//
//	mockPinterestAPI.On("GetPinsBySearch", mock.Anything).Return(&[]Pin{{Id: "123", Url: "http://example.com/image.jpg"}}, nil)
//
//	imageList := []string{"media/test_tag/123.jpg"}
//	tags.UserImageStore.ImageLists = map[int32][]string{123: imageList}
//
//	update := telego.Update{
//		Message: &telego.Message{
//			Chat: telego.Chat{
//				ID: 123,
//			},
//		},
//	}
//
//	sendPhotoParams := tags.handleImage(telego.ID(123), update)
//
//	assert.Equal(t, telego.ID(123), sendPhotoParams.ChatID)
//	assert.NotNil(t, sendPhotoParams.Photo)
//	assert.Equal(t, false, sendPhotoParams.DisableNotification)
//	assert.Equal(t, telego.ModeHTML, sendPhotoParams.ParseMode)
//	assert.Equal(t, "", sendPhotoParams.Caption)
//
//	mockPinterestAPI.AssertExpectations(t)
//}
//
//func TestTagsCommand_NextImageQuery(t *testing.T) {
//	tags := &TagsCommand{
//	}
//
//	imageList := []string{"media/test_tag/123.jpg", "media/test_tag/456.jpg"}
//	tags.UserImageStore.ImageLists = map[int32][]string{123: imageList}
//
//	tags.NextImageQuery(123)
//
//}
