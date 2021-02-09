module gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/login_service

go 1.15

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model => ../common/model

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch => ../common/launch

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/daos/userdao => ../common/daos/userdao

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/dbconnection => ../common/dbconnection

replace gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/response => ../common/response

require (
	github.com/go-chi/chi v1.5.1 // indirect
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/daos/userdao v0.0.0-00010101000000-000000000000
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/dbconnection v0.0.0-00010101000000-000000000000
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/launch v0.0.0-00010101000000-000000000000 // indirect
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/model v0.0.0-00010101000000-000000000000
	gitlab.com/HP-SCDS/Observatorio/2020-2021/uva-devtest/BackEnd/common/response v0.0.0-00010101000000-000000000000
)
