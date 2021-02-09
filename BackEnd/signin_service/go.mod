module gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/signin_service

go 1.15

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model => ../common/model

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch => ../common/launch

require (
	github.com/go-chi/chi v1.5.1 // indirect
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model v0.0.0-00010101000000-000000000000
)
