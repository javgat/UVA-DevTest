import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { UserService } from '@javgat/devtest-api';
import { DataService } from 'src/app/shared/data.service';
import { SessionService } from 'src/app/shared/session.service';
import { NavBarLoggedComponent } from '../nav-bar-logged.component';

@Component({
  selector: 'app-nv-teacher-admin',
  templateUrl: './nv-teacher-admin.component.html',
  styleUrls: ['../../nav-bar.component.css']
})
export class NvTeacherAdminComponent extends NavBarLoggedComponent implements OnInit {

  constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
    super(session, router, data, userS)
  }

  ngOnInit(): void {
  }

}
