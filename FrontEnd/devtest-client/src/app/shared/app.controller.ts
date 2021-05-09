import { Component, OnDestroy } from "@angular/core";
import { Router } from "@angular/router";
import { User, UserService } from "@javgat/devtest-api";
import { Subscription } from "rxjs";
import { Mensaje, SessionLogin, SessionUser } from "./app.model";
import { DataService } from "./data.service";
import { SessionService } from "./session.service";

export class LoggedInController {
    private sessionLogin: SessionLogin
    private sessionUser: SessionUser
    private sessionSubscription: Subscription
    private sessionUserSubscription: Subscription
    private mensaje: Mensaje
    private messageSubscription: Subscription

    constructor(protected session: SessionService, protected router: Router, private data: DataService, protected userS: UserService) {
        this.session.checkStorageSession()
        this.sessionLogin = new SessionLogin(false)
        this.mensaje = new Mensaje()
        this.sessionSubscription = this.session.sessionLogin.subscribe(
            valor => {
                this.sessionLogin = valor
                if (!this.sessionLogin.isLoggedIn()) {
                    this.doActionIsNotLoggedIn()
                }
            }
        )
        this.sessionUser = new SessionUser()
        this.sessionUserSubscription = this.session.sessionUser.subscribe(
            valor => {
                this.sessionUser = valor
                if (this.sessionLogin.isLoggedIn() && this.sessionUser.isEmpty()) {
                    this.downloadUser(true)
                }else{
                    this.doInheritHasUserAction()
                    this.doHasUserAction()
                }
            }
        )

        this.messageSubscription = this.data.mensajeActual.subscribe(
            valor => this.mensaje = valor
        )
    }

    doActionIsNotLoggedIn(){
        this.router.navigate(['/'])
    }

    doInheritHasUserAction(){}
    doHasUserAction(){}

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
                this.session.cambiarUser(new SessionUser(resp.username, resp.email, resp.fullname, resp.rol))
            },
            err => {
                this.handleErrRelog(err, "obtencion del usuario", primera, this.downloadUser, this)
            }
        )
    }

    onDestroy() {
        this.sessionSubscription.unsubscribe()
        this.sessionUserSubscription.unsubscribe()
        this.messageSubscription.unsubscribe();
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

}

export class LoggedInTeacherController extends LoggedInController{

    constructor(session: SessionService, router: Router, data: DataService, userS: UserService) {
        super(session, router, data, userS)
    }

    doInheritHasUserAction(){
        super.doInheritHasUserAction()
        if(this.getSessionUser().isStudent()){
            this.router.navigate(['/'])
        }
    }
}