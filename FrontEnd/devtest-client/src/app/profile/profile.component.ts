import { Component, OnInit } from '@angular/core';
import { ActivatedRoute, Router } from '@angular/router';
import { AuthService, LoginUser, PasswordUpdate, Role, TipoRol, TiporolService, User, UserService, UserUpdate } from '@javgat/devtest-api';
import { Subscription } from 'rxjs';
import { LoggedInController } from '../shared/app.controller';
import { Mensaje, SessionUser, Tipo, Usuario } from '../shared/app.model';
import { DataService } from '../shared/data.service';
import { SessionService } from '../shared/session.service';

@Component({
  selector: 'app-profile',
  templateUrl: './profile.component.html',
  styleUrls: ['./profile.component.css']
})
export class ProfileComponent extends LoggedInController implements OnInit {

  profileUser: Usuario
  private routeSub: Subscription
  id: string
  pUpdate: PasswordUpdate = {
    oldpass: "",
    newpass: "",
  }

  editUser: UserUpdate = {
    username: "",
    email: "",
    fullname: "",
    password: ""
  }
  editRol: Role = {
    rolId: 0,
  }
  editRolChecked: boolean

  newpassR: string
  tiporoles: TipoRol[]
  constructor(session: SessionService, router: Router, private route: ActivatedRoute,
    userS: UserService, data: DataService, private authService: AuthService, private trS?: TiporolService) {
    super(session, router, data, userS)
    this.profileUser = new SessionUser()
    this.id=""
    this.newpassR=""
    this.tiporoles = []
    this.editRolChecked = false
    this.getTipoRoles(true)
    this.routeSub = this.route.params.subscribe(params => {
      this.id = params['id']
      this.borrarMensaje()
      this.doProfileAction()
    });
  }

  doProfileAction(): void{
    this.getProfile(true)
  }

  ngOnInit(): void {
  }

  ngOnDestroy(): void {
    this.routeSub.unsubscribe();
    this.borrarMensaje()
    super.onDestroy()
  }

  getProfile(primera: boolean){
    if (this.id != "") {
      this.getProfileUser(this.id, primera)
    } else{
      let msg = "No se pudo obtener el id del usuario"
      this.cambiarMensaje(new Mensaje(msg, Tipo.ERROR, true))
      console.log(msg)
    }
  }

  getProfileUser(id: string, primera: boolean): void {
    this.userS.getUser(id).subscribe(
      resp => {
        this.profileUser = new Usuario(resp.username, resp.email, resp.fullname, resp.rol, resp.tiporol)
        this.editUser.username = resp.username
        this.editUser.email = resp.email
        this.editUser.fullname = resp.fullname
      },
      err => {
        this.handleErrRelog(err, "obtener datos de perfil de usuario", primera, this.getProfile, this)
      }
    )
  }

  getTipoRoles(primera: boolean): void{
    this.trS?.getTipoRoles().subscribe(
      resp => this.tiporoles = resp,
      err => this.handleErrRelog(err, "obtener tipos de roles", primera, this.getTipoRoles, this)
    )
  }

  changePassSubmit(): void {
    if(this.newpassR==this.pUpdate.newpass){
      this.changePass(true)
    }else{
      this.cambiarMensaje(new Mensaje("Error al actualizar la contraseña: Las contraseñas no coinciden", Tipo.ERROR, true))
    }
  }

  changePass(primera: boolean){
    if (this.id == "") return
    this.userS.putPassword(this.id, this.pUpdate).subscribe(
      resp => {
        console.log("Contraseña cambiada")
        if (this.id == this.getSessionUser().username) {
          var loginUser: LoginUser = {
            loginid: this.getSessionUser().username,
            pass: this.pUpdate.newpass
          }
          this.authService.login(loginUser).subscribe(
            resp => {
              console.log("Sesion JWT recuperada con exito tras cambiar la contraseña")
              this.cambiarMensaje(new Mensaje("Contraseña cambiada con éxito", Tipo.SUCCESS, true))
            },
            err => {
              this.handleShowErr(err, "recuperar sesion JWT tras cambiar la contraseña")
              this.logout()
            }
          )
        }
      },
      err => {
        this.handleErrRelog(err, "cambio de contraseña", primera, this.changePass, this)
      }
    )
  }

  checkPermisosEditarUser(): boolean{
    let us = this.getSessionUser()
    return us.isAdmin() || (us.getUsername()==this.profileUser.getUsername())
  }

  editUserSubmit(){
    this.updateUser(true)
  }

  updateUser(primera: boolean){
    if (this.id == "") return
    this.userS.putUser(this.id, this.editUser).subscribe(
      resp =>{
        this.editUser.password = ""
        this.id = this.editUser.username
        this.cambiarMensaje(new Mensaje("Perfil actualizado con éxito", Tipo.SUCCESS, true))
        this.router.navigate(['/profile', this.id])
        this.getProfile(true)
      },
      err => {
        if(err.status == 409){
          this.cambiarMensaje(new Mensaje("Ya usuario con ese nombre de usuario o correo electrónico", Tipo.ERROR, true))
        }else
          this.handleErrRelog(err, "actualizar datos del usuario", primera, this.updateUser, this)
      }
    )
  }

  checkAdmin(): boolean{
    return this.getSessionUser().isAdmin()
  }

  onSelectRol(rolid: number){
    this.editRolChecked = true
    this.editRol.rolId = rolid
  }

  changeRolSubmit(){
    this.changeRol(true)
  }
  
  changeRol(primera: boolean){
    if (this.id == "") return
    this.userS.putRole(this.id, this.editRol).subscribe(
      resp=>{
        this.cambiarMensaje(new Mensaje("Rol actualizado con éxito", Tipo.SUCCESS, true))
        this.getProfile(true)
      },
      err => this.handleErrRelog(err, "cambiar rol de un usuario", primera, this.changeRol, this)
    )
  }

  checkTeacherOrAdmin(): boolean{
    return this.getSessionUser().isTeacherOrAdmin()
  }

  checkVerBotonesTeacher() : boolean{
    return this.checkPermisosEditarUser() && this.profileUser.isTeacherOrAdmin()
  }

  isValidPass(pass: string): boolean{
    return pass.length>=6
  }

  isDisabledModalEditRolGuardar(): boolean{
    return !this.editRolChecked
  }

  isDisabledModalEditUserGuardar(): boolean{
    return !this.isValidPass(this.editUser.password)
  }

  isDisabledModalEditPassGuardar(): boolean{
    return !(this.isValidPass(this.pUpdate.newpass) && this.isValidPass(this.pUpdate.oldpass) && this.isValidPass(this.newpassR))
  }

  profileNotStudent(): boolean{
    return !this.profileUser.isStudent()
  }

  checkPermisoVerPrivado(): boolean{
    return this.getSessionUser().isAdmin()
  }

}