package service

type IServer interface {
	GetConfig()
	Report()
}

type Server struct {
	BaseUrl   string
	ConfigUrl string
	ReportUrl string
}

var s *Server

func New(config map[string]string){
	s = &Server{
		config["base"],
		config["config"],
		config["report"],
	}
}

func GetServer() *Server{
	if s == nil {
		panic("server error")
	}
	return s
}

func (s *Server) GetConfig(projectName string) {
	// todo 请求配置文件
}

func (s *Server) Report(projectName string) {
	// todo 上报文件
}
