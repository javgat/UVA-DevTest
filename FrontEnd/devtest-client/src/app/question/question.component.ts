import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { Option, Question, QuestionService, UserService } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, Tipo } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';


@Component({
  selector: 'app-question',
  templateUrl: './question.component.html',
  styleUrls: ['./question.component.css']
})
export class QuestionComponent extends LoggedInController implements OnInit {

  question: Question
  routeSub: Subscription
  id: number
  tipoPrint: string
  opciones: Option[]
  op: Option | undefined
  nuevaOpcion: Option
  deleteIndex: number
  constructor(session: SessionService, router: Router, data: DataService, userS: UserService, private qS: QuestionService, private route: ActivatedRoute) {
    super(session, router, data, userS)
    this.opciones = []
    this.deleteIndex = -1
    this.question = {
      title: "",
      question: "",
      estimatedTime: 0,
      autoCorrect: false,
      editable: true,
      username: "",
      tipoPregunta: Question.TipoPreguntaEnum.String
    }
    this.nuevaOpcion = {
      correcta: false,
      texto: ""
    }
    this.tipoPrint = ""
    this.setTipoPrint()
    this.id = 0
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      this.getPregunta(true)
    });
  }

  ngOnInit(): void {
  }

  setTipoPrint(){
    switch(this.question.tipoPregunta){
      case Question.TipoPreguntaEnum.String:
        this.tipoPrint = "Texto"
        break
      case Question.TipoPreguntaEnum.Codigo:
        this.tipoPrint = "Código"
        break
      case Question.TipoPreguntaEnum.Opciones:
        if(this.question.eleccionUnica)
          this.tipoPrint = "Tipo test de respuesta única"
        else
          this.tipoPrint = "Tipo test de respuesta múltiple"
        break
      default:
        this.tipoPrint = ""
    }
  }

  getPregunta(primera: boolean) {
    this.qS.getQuestion(this.id).subscribe(
      resp => {
        this.question = resp
        this.setTipoPrint()
        if(this.question.tipoPregunta == Question.TipoPreguntaEnum.Opciones){
          this.getOptions(true)
        }
      },
      err => this.handleErrRelog(err, "obtener pregunta", primera, this.getPregunta, this)
    )
  }

  getOptions(primera: boolean){
    this.qS.getOptionsFromQuestion(this.id).subscribe(
      resp=>{
        this.opciones = resp
      },
      err => this.handleErrRelog(err, "obtener opciones de respuesta de pregunta", primera, this.getOptions, this)
    )
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe()
    super.onDestroy()
  }

  checkPermisosEdicion() : boolean{
    return true
  }

  changeCorrectaOpc(op: Option){
    this.op = op 
    this.changeCorrecta(true)
  }

  changeCorrecta(primera: boolean){
    if(this.op == undefined) return
    this.op.correcta = !this.op.correcta
    if(this.op.indice==undefined){
      this.cambiarMensaje(new Mensaje("Error, falta el campo indice en la opcion", Tipo.ERROR, true))
      return
    }
    this.qS.putOption(this.id, this.op.indice, this.op).subscribe(
      resp =>{},
      err => this.handleErrRelog(err, "cambiar campo de correccion de una opcion", primera, this.changeCorrecta, this)
    )
  }

  addOptionSubmit(){
    //this.nuevaOpcion.preguntaid = this.id
    this.addOption(true)
  }

  addOption(primera: boolean){
    this.qS.postOption(this.id, this.nuevaOpcion).subscribe(
      resp=> this.getOptions(true),
      err=> this.handleErrRelog(err, "añadir nueva opcion de respuesta", primera, this.addOption, this)
    )
  }
  
  deleteOpc(indice: number | undefined){
    if(indice!=undefined)
      this.deleteIndex = indice
    this.deleteOpcion(true)
  }

  deleteOpcion(primera: boolean){
    this.qS.deleteOption(this.id, this.deleteIndex).subscribe(
      resp=> this.getOptions(true),
      err=> this.handleErrRelog(err, "eliminar una opcion", primera, this.deleteOpcion, this)
    )
  }
}
