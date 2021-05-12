import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Question, QuestionService, TagService, TeamService, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-list-questions',
  templateUrl: './list-questions.component.html',
  styleUrls: ['./list-questions.component.css']
})
export class ListQuestionsComponent extends LoggedInController implements OnInit {

  newSearchTag: string
  questions: Question[]
  searchTags: string[][]
  includeNonEdit: boolean
  orOperation: boolean
  likeTitle: string | undefined
  editLikeTitle: string
  hideSwitchInclude: boolean
  scrollable: boolean
  selectAddQuestion: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService,
     protected qS: QuestionService, protected tagS: TagService, protected teamS?: TeamService) {
    super(session, router, data, userS)
    this.includeNonEdit = false
    this.searchTags = []
    this.questions = []
    this.newSearchTag = ""
    this.orOperation = true
    this.editLikeTitle = ""
    this.hideSwitchInclude = false
    this.scrollable = false
    this.selectAddQuestion = false
    this.getQuestionsFilters()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  // esta funcion se tiene que sobreescribir
  getQuestionsInclude(primera: boolean) {
    this.qS.getQuestions(this.searchTags, this.likeTitle).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas", primera, this.getQuestionsInclude, this)
    )
  }

  // esta funcion se tiene que sobreescribir
  getQuestionsEdit(primera: boolean) {
    this.qS.getEditQuestions(this.searchTags, this.likeTitle).subscribe(
      resp => this.questions = resp,
      err => this.handleErrRelog(err, "obtener preguntas no publicadas", primera, this.getQuestionsEdit, this)
    )
  }

  selectQuestion(id: number | undefined){}

  getQuestionsFilters() {
    if (this.includeNonEdit) {
      this.getQuestionsInclude(true)
    } else {
      this.getQuestionsEdit(true)
    }
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined) {
    return tipoPrint(tipo, eleUni)
  }

  changeFlexInclude() {
    this.includeNonEdit = !this.includeNonEdit
    this.getQuestionsFilters()
  }

  addSearchTagOr() {
    var orTags = [this.newSearchTag]
    this.searchTags.push(orTags)
    this.newSearchTag = ""
  }

  addSearchTagAnd() {
    if(this.searchTags.length==0) this.searchTags.push([])
    this.searchTags[0].push(this.newSearchTag)
    this.newSearchTag = ""
  }

  addSearchTagSubmit() {
    if (this.orOperation) {
      this.addSearchTagOr()
    } else {
      this.addSearchTagAnd()
    }
    this.getQuestionsFilters()
  }

  deleteSearchTag(deleteTag: string) {
    this.searchTags.forEach((element, arrIndex) => {
      element.forEach((subElement, index) => {
        if (subElement == deleteTag) element.splice(index, 1)
      })
      if(element.length==0) this.searchTags.splice(arrIndex, 1)
    });
  }

  deleteSearchTagClick(deleteTag: string) {
    this.deleteSearchTag(deleteTag)
    this.getQuestionsFilters()
  }

  changeToAndTags() {
    if(this.orOperation){
      this.orOperation = false
      var andTags: string[] = []
      this.searchTags.forEach(element => {
        element.forEach(subElem => {
          andTags.push(subElem)
        })
      })
      this.searchTags = [andTags]
    }
  }

  changeToOrTags() {
    if(!this.orOperation){
      this.orOperation = true
      var orTags: string[][] = []
      this.searchTags.forEach(element => {
        element.forEach(subElem => {
          orTags.push([subElem])
        })
      })
      this.searchTags = orTags
    }
  }

  swapOrAndTags(){
    if(this.orOperation){
      this.changeToAndTags()
    }else{
      this.changeToOrTags()
    }
    this.getQuestionsFilters()
  }

  clickSearchTitle(){
    this.likeTitle = this.editLikeTitle
    this.getQuestionsFilters()
  }

  clickBorrarTitle(){
    this.likeTitle = undefined
    this.editLikeTitle = ""
    this.getQuestionsFilters()
  }

  newTabQuestion(id: number | undefined){
    const url = this.router.serializeUrl(
      this.router.createUrlTree(['/q', id])
    );
  
    window.open(url, '_blank');
  }

}
