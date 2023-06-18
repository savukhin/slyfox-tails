package api

type LoginUserDTO struct {
	Username string `validate:"required" json:"username"`
	Password string `validate:"required" json:"password"`
}

type RegisterUserDTO struct {
	Username       string `validate:"required,min=3,max=32" json:"username"`
	Email          string `validate:"required,email,min=6,max=32" json:"email"`
	Password       string `validate:"required,min=3,max=40" json:"password"`
	PasswordRepeat string `validate:"required,min=3,max=40" json:"password_repeat"`
}

type CreateProjectDTO struct {
	Title string `validate:"required,min=3,max=32" json:"title"`
}

type CreateJobDTO struct {
	Title     string `validate:"required,min=3,max=32" json:"title"`
	ProjectID uint64 `validate:"required,number" json:"project_id"`
}

type UpdateJobDTO struct {
	Title string `validate:"required,min=3,max=32" json:"title"`
}

type CreateStageDTO struct {
	Title string `validate:"required,min=3,max=32" json:"title"`
	JobID uint64 `validate:"required,number" json:"job_id"`
}

type UpdatedStageDTO struct {
	Title       string `validate:"min=3,max=32" json:"title"`
	StartedAtMs uint64 `validate:"numeric" json:"started_at_ms"`
}

type CreatePointDTO struct {
	Title  string   `validate:"required,min=3,max=32" json:"title"`
	Stages []uint64 `validate:"" json:"stage_ids,omitempty"`
}

type UpdatePointDTO struct {
	Title  string   `validate:"min=3,max=32" json:"title,omitempty"`
	Stages []uint64 `validate:"" json:"stage_ids,omitempty"`
}
