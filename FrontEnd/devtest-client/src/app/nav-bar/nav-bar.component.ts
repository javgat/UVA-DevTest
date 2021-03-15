import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { SessionLogin, SessionUser } from '../shared/app.model';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-nav-bar',
  templateUrl: './nav-bar.component.html',
  styleUrls: ['./nav-bar.component.css']
})
export class NavBarComponent implements OnInit {


  private sessionLogin : SessionLogin
  sessionUser : SessionUser
  private sessionSubscription : Subscription
  private sessionUserSubscription : Subscription
  
  constructor(private session: SessionService, private router: Router, private userService : UserService) {
    this.session.checkStorageSession()
    this.sessionLogin = new SessionLogin(false)
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionLogin = valor
        if(!this.sessionLogin.logged){
          this.router.navigate(['/'])
        }
      }
    )
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor => this.sessionUser = valor
    )
  }

  ngOnInit(): void {
    if (this.sessionUser.isEmpty){
      this.getUser()
    }
  }

  ngOnDestroy(): void {
    this.sessionSubscription.unsubscribe();
    this.sessionUserSubscription.unsubscribe();
  }

  logout(){
    this.session.logout()
  }


  getUser(){
    this.userService.getUser(this.sessionLogin.userid as string).subscribe(
      resp => {
        this.session.cambiarSession(new SessionLogin(true, this.sessionLogin.userid))
        this.session.cambiarUser(new SessionUser(resp.username, resp.email, resp.fullname, resp.type))
      },
      err => {
        console.log("No se pudo obtener el usuario")
      }
    )
  }

}
