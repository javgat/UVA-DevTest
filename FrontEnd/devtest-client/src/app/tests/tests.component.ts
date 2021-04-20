import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { TagService, Test, TestService, UserService } from '@javgat/devtest-api';
import { LoggedInTeacherController } from '../shared/app.controller';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-tests',
  templateUrl: './tests.component.html',
  styleUrls: ['./tests.component.css']
})
export class TestsComponent extends LoggedInTeacherController implements OnInit {

  tests: Test[]
  searchTag: string
  includeNonEdit: boolean
  tagPressed: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private tS: TestService, private tagS: TagService) {
    super(session, router, data, userS)
    this.includeNonEdit = false
    this.tagPressed = false
    this.searchTag = ""
    this.tests = []
    this.getTests(true)
  }

  ngOnInit(): void {
  }


  ngOnDestroy(): void {
    super.onDestroy()
  }

  getTests(primera: boolean) {
    if (this.includeNonEdit) {
      this.tS.getPublicTests().subscribe(
        resp => this.tests = resp,
        err => this.handleErrRelog(err, "obtener tests", primera, this.getTests, this)
      )
    } else {
      this.tS.getPublicEditTests().subscribe(
        resp => this.tests = resp,
        err => this.handleErrRelog(err, "obtener tests editables", primera, this.getTests, this)
      )
    }
  }

  getTestsWithTagSubmit() {
    this.tagPressed=true
    this.getTestsWithTag(true)
  }

  getTestsWithTag(primera: boolean) {
    if (this.includeNonEdit) {
      this.tagS.getTestsFromTag(this.searchTag).subscribe(
        resp => this.tests = resp,
        err => this.handleErrRelog(err, "obtener tests por etiqueta", primera, this.getTestsWithTag, this)
      )
    } else {
      this.tagS.getEditTestsFromTag(this.searchTag).subscribe(
        resp => this.tests = resp,
        err => this.handleErrRelog(err, "obtener preguntas no publicadas por etiqueta", primera, this.getTestsWithTag, this)
      )
    }
  }

  changeFlexInclude() {
    this.includeNonEdit = !this.includeNonEdit
    if(this.tagPressed){
      this.getTestsWithTag(true)
    }else{
      this.getTests(true)
    }
  }

}
