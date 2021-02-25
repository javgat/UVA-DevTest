import { Component, OnInit } from '@angular/core';
import { AuthService, LoginUser } from '@javgat/devtest-api'

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  loginUserEmpty = {
    loginid: "",
    pass: ""
  }

  loginUser = this.loginUserEmpty as LoginUser

  constructor(private authService : AuthService) {
    
  }

  ngOnInit(): void {
  }

  login(lu : LoginUser){
    this.authService.login(lu).subscribe(
      resp => {
        console.log("Inicio sesion con exito")
      },
      err =>{
        console.log("Error al iniciar sesion")
      }
    )
  }

  loginSubmit(){
    this.login(this.loginUser)
  }

}
