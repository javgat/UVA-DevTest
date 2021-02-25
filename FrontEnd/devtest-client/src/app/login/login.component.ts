import { Component, OnInit } from '@angular/core';
import { AuthService, LoginUser } from '@javgat/devtest-api'

@Component({
  selector: 'app-login',
  templateUrl: './login.component.html',
  styleUrls: ['./login.component.css']
})
export class LoginComponent implements OnInit {

  constructor(authService : AuthService) {
    let loginUser: LoginUser = {
      loginid: "Carlos",
      pass: "carlospass"
    }

    authService.login(loginUser).subscribe(console.log)
  }

  ngOnInit(): void {
  }

}
