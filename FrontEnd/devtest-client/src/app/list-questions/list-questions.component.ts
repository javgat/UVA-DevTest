import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Question, QuestionService, Tag, TagService, TeamService, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { EnumOrderBy, Mensaje, Tipo, tipoPrint } from '../shared/app.model';
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
  orderBy: EnumOrderBy
  limit: number
  offset: number
  autotags: Tag[]
  mensajeListaVacia: string
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
    this.orderBy = EnumOrderBy.newDate
    this.limit = 20
    this.offset = 0
    this.autotags = []
    this.mensajeListaVacia = "¡Vaya! Parece que no hay ninguna pregunta para mostrar"
    this.getQuestionsFilters()
    this.changeGetAutoTags()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  // esta funcion se tiene que sobreescribir
  getQuestionsInclude(primera: boolean) {
    this.qS.getQuestions(this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas", primera, this.getQuestionsInclude, this)
    )
  }

  // esta funcion se tiene que sobreescribir
  getQuestionsEdit(primera: boolean) {
    this.qS.getEditQuestions(this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.saveQuestions(resp),
      err => this.handleErrRelog(err, "obtener preguntas no publicadas", primera, this.getQuestionsEdit, this)
    )
  }

  selectQuestion(id: number | undefined){}

  saveQuestions(resp: Question[]){
    this.questions = resp
    this.borrarMensaje()
  }

  getQuestionsFilters() {
    this.cambiarMensaje(new Mensaje("Descargando preguntas... ", Tipo.DOWNLOADING, true))
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

  clickOrderByTiempo(){
    if(this.orderBy == EnumOrderBy.moreTime){
      this.orderBy = EnumOrderBy.lessTime
    }else if(this.orderBy == EnumOrderBy.lessTime){
      this.orderBy = EnumOrderBy.newDate
    }else{
      this.orderBy = EnumOrderBy.moreTime
    }
    this.getQuestionsFilters()
  }

  clickOrderByFavoritos(){
    if(this.orderBy == EnumOrderBy.moreFav){
      this.orderBy = EnumOrderBy.lessFav
    }else if(this.orderBy == EnumOrderBy.lessFav){
      this.orderBy = EnumOrderBy.newDate
    }else{
      this.orderBy = EnumOrderBy.moreFav
    }
    this.getQuestionsFilters()
  }

  isMoreTimeSelected(): boolean{
    return this.orderBy == EnumOrderBy.moreTime
  }

  isLessTimeSelected(): boolean{
    return this.orderBy == EnumOrderBy.lessTime
  }

  isMoreFavSelected(): boolean{
    return this.orderBy == EnumOrderBy.moreFav
  }

  isLessFavSelected(): boolean{
    return this.orderBy == EnumOrderBy.lessFav
  }

  getCurrentPage(): number{
    return (this.offset/this.limit)+1
  }

  hasNextPage(): boolean{
    return this.questions.length == this.limit
  }

  clickPreviousPage(){
    this.offset = this.offset-this.limit
    this.getQuestionsFilters()
  }

  clickNextPage(){
    this.offset = this.offset-this.limit
    this.getQuestionsFilters()
  }

  clickFirstPage(){
    this.offset=0
    this.getQuestionsFilters()
  }

  clickCurrentPage(){
    this.getQuestionsFilters()
  }


  changeGetAutoTags(){
    this.getAutoTags(true)
  }

  getAutoTags(primera: boolean){
    this.tagS.getTags(this.newSearchTag ,"moreQuestion", 20).subscribe(
      resp=>{
        this.autotags=resp
      },
      err => this.handleErrRelog(err, "obtener tags de preguntas mas comunes", primera, this.getAutoTags, this)
    )
  }

}
