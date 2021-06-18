import { Component, OnInit } from '@angular/core';
import { Router } from '@angular/router';
import { Tag, TagService, TeamService, Test, TestService, UserService } from '@javgat/devtest-api';
import { LoggedInController } from '../shared/app.controller';
import { EnumOrderBy, Mensaje, Tipo, tipoPrint } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';
@Component({
  selector: 'app-list-tests',
  templateUrl: './list-tests.component.html',
  styleUrls: ['./list-tests.component.css']
})
export class ListTestsComponent extends LoggedInController implements OnInit {

  newSearchTag: string
  tests: Test[]
  searchTags: string[][]
  includeNonEdit: boolean
  orOperation: boolean
  likeTitle: string | undefined
  editLikeTitle: string
  hideSwitchInclude: boolean
  includeLabel: string
  arePublished: boolean
  orderBy: EnumOrderBy
  limit: number
  offset: number
  autotags: Tag[]
  mensajeListaVacia: string
  searchActive: boolean
  filterActive: boolean
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, protected tS: TestService,
     protected tagS: TagService, protected teamS?: TeamService) {
    super(session, router, data, userS)
    this.includeNonEdit = false
    this.searchTags = []
    this.tests = []
    this.newSearchTag = ""
    this.orOperation = true
    this.editLikeTitle = ""
    this.hideSwitchInclude = true
    this.arePublished = false
    this.includeLabel = "Incluir tests publicados"
    this.orderBy = EnumOrderBy.newDate
    this.limit = 20
    this.offset = 0
    this.autotags = []
    this.searchActive = false
    this.filterActive = false
    this.mensajeListaVacia = "¡Vaya! Parece que no hay ningún test para mostrar"
    this.getTestsFilters()
    this.changeGetAutoTags()
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    super.onDestroy()
  }

  // esta funcion se tiene que sobreescribir
  getTestsInclude(primera: boolean) {
    this.getTestsEdit(primera)
  }

  // esta funcion se tiene que sobreescribir
  getTestsEdit(primera: boolean) {
    this.tS.getPublicEditTests(this.searchTags, this.likeTitle, this.orderBy, this.limit, this.offset).subscribe(
      resp => this.saveTests(resp),
      err => this.handleErrRelog(err, "obtener preguntas no publicadas", primera, this.getTestsEdit, this)
    )
  }

  saveTests(resp: Test[]){
    this.tests = resp
    this.borrarMensaje()
  }

  getTestsFilters() {
    this.cambiarMensaje(new Mensaje("Descargando tests... ", Tipo.DOWNLOADING, true))
    if (this.includeNonEdit) {
      this.getTestsInclude(true)
    } else {
      this.getTestsEdit(true)
    }
  }

  tipoPrint(tipo: string, eleUni: boolean | undefined) {
    return tipoPrint(tipo, eleUni)
  }

  changeFlexInclude() {
    this.includeNonEdit = !this.includeNonEdit
    this.getTestsFilters()
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
    this.getTestsFilters()
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
    this.getTestsFilters()
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

  clickOrTags(){
    this.changeToOrTags()
    this.getTestsFilters()
  }

  clickAndTags(){
    this.changeToAndTags()
    this.getTestsFilters()
  }

  clickSearchTitle(){
    this.likeTitle = this.editLikeTitle
    this.getTestsFilters()
  }

  clickBorrarTitle(){
    this.likeTitle = undefined
    this.editLikeTitle = ""
    this.getTestsFilters()
  }

  clickOrderByTiempo(){
    if(this.orderBy == EnumOrderBy.moreTime){
      this.orderBy = EnumOrderBy.lessTime
    }else if(this.orderBy == EnumOrderBy.lessTime){
      this.orderBy = EnumOrderBy.newDate
    }else{
      this.orderBy = EnumOrderBy.moreTime
    }
    this.getTestsFilters()
  }

  clickOrderByFavoritos(){
    if(this.orderBy == EnumOrderBy.moreFav){
      this.orderBy = EnumOrderBy.lessFav
    }else if(this.orderBy == EnumOrderBy.lessFav){
      this.orderBy = EnumOrderBy.newDate
    }else{
      this.orderBy = EnumOrderBy.moreFav
    }
    this.getTestsFilters()
  }

  clickOrderByFechaPub(){
    if(this.orderBy == EnumOrderBy.oldDate){
      this.orderBy = EnumOrderBy.newDate
    }else{
      this.orderBy = EnumOrderBy.oldDate
    }
    this.getTestsFilters()
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

  isNewDateSelected(): boolean{
    return this.orderBy == EnumOrderBy.newDate
  }

  isOldDateSelected(): boolean{
    return this.orderBy == EnumOrderBy.oldDate
  }

  getCurrentPage(): number{
    return (this.offset/this.limit)+1
  }

  hasNextPage(): boolean{
    return this.tests.length == this.limit
  }

  clickPreviousPage(){
    this.offset = this.offset-this.limit
    this.getTestsFilters()
  }

  clickNextPage(){
    this.offset = this.offset-this.limit
    this.getTestsFilters()
  }

  clickFirstPage(){
    this.offset=0
    this.getTestsFilters()
  }

  clickCurrentPage(){
    this.getTestsFilters()
  }

  changeGetAutoTags(){
    this.getAutoTags(true)
  }

  getAutoTags(primera: boolean){
    this.tagS.getTags(this.newSearchTag ,"moreTest", 20).subscribe(
      resp=>{
        this.autotags=resp
      },
      err => this.handleErrRelog(err, "obtener tags de tests mas comunes", primera, this.getAutoTags, this)
    )
  }
  
  printDate(d: Date | undefined): string{
    if(d==undefined) return ""
    var date = new Date(d)
    return date.toLocaleString()
  }

  activateSearch(){
    this.searchActive=true
  }

  isSearchActive(): boolean{
    return this.searchActive
  }

  activateFilter(){
    this.filterActive=true
  }

  isFilterActive(): boolean{
    return this.filterActive
  }
}