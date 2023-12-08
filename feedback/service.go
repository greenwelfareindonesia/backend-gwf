package feedback

type Service interface {
	GetAllFeedback(input int) ([]Feedback, error)
	GetFeedbackByID(id int) (Feedback, error)
	CreateFeedback(feedback FeedbackInput) (Feedback, error)
	DeleteFeedback(ID int) (Feedback, error)
	// UpdateFeedback(getIdFeedback FeedbackID, input FeedbackInput, FileName string) (Feedback, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) DeleteFeedback(ID int) (Feedback, error) {
	feedbacks, err := s.repository.FindById(ID)
	if err != nil {
		return feedbacks, err
	}

	feedback, err := s.repository.DeleteFeedback(feedbacks)
	if err != nil {
		return feedback, err
	}
	return feedback, nil
}


func (s *service) GetAllFeedback(input int) ([]Feedback, error) {
	feedbacks, err := s.repository.FindAll()
	if err != nil {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (s *service) GetFeedbackByID(id int) (Feedback, error) {
	feedbacks, err := s.repository.FindById(id)
	if err != nil {
		return feedbacks, err
	}
	if feedbacks.ID == 0 {
		return feedbacks, err
	}
	return feedbacks, nil
}

func (s *service) CreateFeedback(feedback FeedbackInput) (Feedback, error) {
	newFeedback := Feedback{}

	newFeedback.Email = feedback.Email
	newFeedback.Text = feedback.Text

	saveFeedback, err := s.repository.Create(newFeedback)

	if err != nil {
		return saveFeedback, err
	}
	return saveFeedback, nil
}
