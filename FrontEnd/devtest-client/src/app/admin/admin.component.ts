import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { User, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { SessionUser } from '../shared/app.model';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-admin',
  templateUrl: './admin.component.html',
  styleUrls: ['./admin.component.css']
})
export class AdminComponent implements OnInit {

  users : User[]
  sessionUser : SessionUser

  private sessionUserSubscription : Subscription

  constructor(private session: SessionService, private userService: UserService,
    private router : Router) {
    this.users = []
    this.sessionUser = new SessionUser()
    this.sessionUserSubscription = this.session.sessionUser.subscribe(
      valor =>{
        this.sessionUser = valor
        this.getUsers()
      }
    )
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    this.sessionUserSubscription.unsubscribe()
  }

  getUsers(){
    this.userService.getUsers().subscribe(
      resp => {
        this.users = resp
      },
      err =>{
        console.log("Error, no tienes permisos para acceder a los usuarios "+err.status)
        // Esta comentado para mostrar que no estan los usuarios
        //this.router.navigate(['/'])
      }
    )
  }

}
