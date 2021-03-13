import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { MainComponent } from '../main.component';

import { UserService} from '@javgat/devtest-api'

@Component({
  selector: 'app-logged-in',
  templateUrl: './logged-in.component.html',
  styleUrls: ['./logged-in.component.css']
})
export class LoggedInComponent extends MainComponent implements OnInit {

  constructor(datos: DataService, session: SessionService, private userService : UserService){
    super(datos, session);
  }

  ngOnInit(): void {
    this.getUser()
  }

  getUser(){
    this.userService.getUser(this.sessionActual.userid as string).subscribe(
      resp => {
        console.log(resp.username)
      },
      err => {
        console.log("No se pudo obtener el usuario")
      }
    )
  }

}
