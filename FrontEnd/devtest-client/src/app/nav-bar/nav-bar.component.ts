import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
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
  isAdmin : boolean
  
  constructor(private session: SessionService, private router: Router, private userService : UserService) {
    this.session.checkStorageSession()
    this.sessionLogin = new SessionLogin(false)
    this.sessionSubscription = this.session.sessionLogin.subscribe(
      valor => {
        this.sessionLogin = valor
        if(!this.sessionLogin.isLoggedIn()){
          this.router.navigate(['/'])
        }
      }
    )
    this.isAdmin = false
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor =>{
        this.sessionUser = valor
        this.isAdmin = (valor.getRol() == User.RolEnum.Administrador)
      }
    )
  }

  ngOnInit(): void {
    if (this.sessionUser.isEmpty()){
      this.getUser(true)
    }
  }

  ngOnDestroy(): void {
    this.sessionSubscription.unsubscribe();
    this.sessionUserSubscription.unsubscribe();
  }

  logout(){
    this.session.logout()
  }

  getUser(primera: boolean){
    this.userService.getUser(this.sessionLogin.getUserUsername() as string).subscribe(
      resp => {
        this.session.cambiarSession(new SessionLogin(true, this.sessionLogin.getUserUsername()))
        this.session.cambiarUser(new SessionUser(resp.username, resp.email, resp.fullname, resp.rol))
      },
      err => {
        this.session.handleErrRelog(err, "obtencion del usuario", primera, this.getUser, this)        
      }
    )
  }

}
