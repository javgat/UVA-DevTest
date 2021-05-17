import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { LoggedInController } from 'src/app/shared/app.controller';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';

@Component({
  selector: 'app-lgi-student',
  templateUrl: './lgi-student.component.html',
  styleUrls: ['./lgi-student.component.css']
})
export class LgiStudentComponent extends LoggedInController implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void{
    super.onDestroy()
  }

}
