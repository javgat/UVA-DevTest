import { Component, OnInit } from '@angular/core';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { MainComponent } from '../main.component';

import { UserService} from '@javgat/devtest-api'
import { Router } from '@angular/router';

@Component({
  selector: 'app-logged-in',
  templateUrl: './logged-in.component.html',
  styleUrls: ['./logged-in.component.css']
})
export class LoggedInComponent extends MainComponent implements OnInit {

  constructor(session: SessionService, router: Router, datos: DataService, userService : UserService){
    super(session, router, datos, userService);
  }

  ngOnInit(): void {
  }


}
