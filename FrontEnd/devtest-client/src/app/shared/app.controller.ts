import { Component, OnDestroy } from "@angular/core";
import { Router } from "@angular/router";
import { TipoRol, User, UserService } from "@javgat/devtest-api";
import { Subscription } from "rxjs";
import { Mensaje, SessionLogin, SessionUser } from "./app.model";
import { DataService } from "./data.service";
import { SessionService } from "./session.service";

export abstract class LoggedInController {
    private sessionLogin: SessionLogin
    private sessionUser: SessionUser
    private sessionTipoRoles: TipoRol[]
    private sessionSubscription: Subscription
    private sessionUserSubscription: Subscription
    private sessionTipoRolesSubscription: Subscription
    private mensaje: Mensaje
    private messageSubscription: Subscription

    private sessionLoginWaited: boolean
    private sessionUserWaited: boolean

    constructor(protected session: SessionService, protected router: Router, private data: DataService, protected userS: UserService) {
        this.session.checkStorageSession()
        this.sessionLogin = new SessionLogin(false)
        this.sessionUser = new SessionUser()
        this.mensaje = new Mensaje()
        this.sessionTipoRoles = []
        this.sessionLoginWaited = false
        this.sessionUserWaited = false
        this.sessionTipoRolesSubscription = this.session.sessionTipoRoles.subscribe(
            valor => {
                this.sessionTipoRoles = valor
                if(this.sessionUserWaited){
                    this.doUserSubscriptionAction()
                }
                if(this.sessionLoginWaited){
                    this.doSessionLoginSubscriptionAction()
                }
            }
        )
        this.sessionSubscription = this.session.sessionLogin.subscribe(
            valor => {
                this.sessionLogin = valor
                this.doSessionLoginSubscriptionAction()
            }
        )
        this.sessionUserSubscription = this.session.sessionUser.subscribe(
            valor => {
                this.sessionUser = valor
                this.doUserSubscriptionAction()
            }
        )
        this.messageSubscription = this.data.mensajeActual.subscribe(
            valor => this.mensaje = valor
        )
    }

    doSessionLoginSubscriptionAction() {
        if (this.sessionTipoRoles.length > 0) {
            if(!this.sessionLogin.isLoggedIn()){
                this.doActionIsNotLoggedIn()
            }
        }else{
            this.sessionLoginWaited = true
        }
    }

    doUserSubscriptionAction() {
        if (this.sessionLogin.isLoggedIn() && this.sessionUser.isEmpty()) {
            this.downloadUser(true)
        } else if (this.sessionTipoRoles.length > 0) {
            this.doInheritHasUserAction()
            this.doHasUserAction()
            this.doActionKnowTipoRol()
        } else{
            this.sessionUserWaited = true
        }
    }

    doActionIsNotLoggedIn() {
        this.doActionKnowTipoRol()
    }

    doInheritHasUserAction() { }
    doHasUserAction() {
    }

    doActionKnowTipoRol() {
        if (!this.hasPermissions()) {
            this.redirectNotAllowed()
        }
    }

    // Importante para sobreescribir esta
    // Defines las condiciones que se tienen que cumplir para que tengas permisos
    hasPermissions(): boolean {
        return this.sessionLogin.isLoggedIn()
    }

    // Redirige cuando no tienes permisos
    redirectNotAllowed() {
        this.router.navigate(['/'])
    }

    getSessionLogin(): SessionLogin {
        return this.sessionLogin
    }

    logout() {
        this.session.logout()
    }

    getSessionUser(): SessionUser {
        return this.sessionUser
    }

    downloadUser(primera: boolean) {
        this.userS.getUser(this.sessionLogin.getUserUsername() as string).subscribe(
            resp => {
                this.session.cambiarUser(new SessionUser(resp.username, resp.email, resp.fullname, resp.rol, resp.tiporol))
            },
            err => {
                this.handleErrRelog(err, "obtencion del usuario", primera, this.downloadUser, this)
            }
        )
    }

    getSessionTipoRoles(): TipoRol[] {
        return this.sessionTipoRoles
    }

    onDestroy() {
        this.sessionSubscription.unsubscribe()
        this.sessionUserSubscription.unsubscribe()
        this.messageSubscription.unsubscribe();
        this.sessionTipoRolesSubscription.unsubscribe()
    }

    getMensaje(): Mensaje {
        return this.mensaje
    }

    cambiarMensaje(m: Mensaje) {
        this.data.cambiarMensaje(m)
    }

    borrarMensaje() {
        this.data.borrarMensaje()
    }

    handleShowErr(err: any, action: string) {
        this.data.handleShowErr(err, action)
    }

    handleErrRelog<T>(err: any, action: string, primera: boolean, callbackFn: (this: T, prim: boolean) => void, that: T): void {
        this.session.handleErrRelog(err, action, primera, callbackFn, that)
    }

    reLoad() {
        this.router.navigate([this.router.url])
    }

    isLoggedIn(): boolean {
        return this.getSessionLogin().isLoggedIn()
    }

    getTipoRolActual(): TipoRol | undefined {
        if (this.sessionLogin.isLoggedIn()) {
            return this.sessionTipoRoles.filter(tipo => tipo.nombre == this.sessionUser.getTipoRol())[0]
        }
        return this.sessionTipoRoles.filter(tipo => tipo.rolBase == TipoRol.RolBaseEnum.NoRegistrado)[0]
    }

    canVerPTests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.verPTests)
            return true
        return false
    }

    canVerETests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.verETests)
            return true
        return false
    }

    canVerEQuestions(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.verEQuestions)
            return true
        return false
    }

    canVerPQuestions(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.verPQuestions)
            return true
        return false
    }

    canVerAnswers(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.verAnswers)
            return true
        return false
    }

    canChangeRoles(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.changeRoles)
            return true
        return false
    }

    canTenerTeams(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.tenerTeams)
            return true
        return false
    }

    canTenerEQuestions(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.tenerEQuestions)
            return true
        return false
    }

    canTenerETests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.tenerETests)
            return true
        return false
    }

    canTenerPTests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.tenerPTests)
            return true
        return false
    }

    canAdminPTests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminPTests)
            return true
        return false
    }

    canAdminETests(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminETests)
            return true
        return false
    }

    canAdminEQuestions(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminEQuestions)
            return true
        return false
    }

    canAdminAnswers(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminAnswers)
            return true
        return false
    }

    canAdminUsers(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminUsers)
            return true
        return false
    }

    canAdminTeams(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminTeams)
            return true
        return false
    }

    canAdminConfiguration(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminConfiguration)
            return true
        return false
    }

    canAdminPermissions(): boolean {
        let rol = this.getTipoRolActual()
        if (rol != undefined && rol.adminPermissions)
            return true
        return false
    }

}

export class LoggedInTeacherController extends LoggedInController {

    constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
        super(session, router, data, userS)
    }

    doInheritHasUserAction() {
        super.doInheritHasUserAction()
        if (this.getSessionUser().isStudent()) {
            this.router.navigate(['/'])
        }
    }
}