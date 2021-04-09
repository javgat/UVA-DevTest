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

    isLoggedIn: boolean
    isAdmin: boolean
    isTeacher: boolean
    isStudent: boolean
    constructor(protected session: SessionService, protected router: Router, private data: DataService, protected userS: UserService) {
        this.isLoggedIn = this.isAdmin = this.isTeacher = this.isStudent = false
        this.session.checkStorageSession()
        this.sessionLogin = new SessionLogin(false)
        this.mensaje = new Mensaje()
        this.sessionSubscription = this.session.sessionLogin.subscribe(
            valor => {
                this.sessionLogin = valor
                this.isLoggedIn = valor.isLoggedIn()
                if (!this.sessionLogin.isLoggedIn()) {
                    this.router.navigate(['/'])
                }
            }
        )
        this.sessionUser = new SessionUser()
        this.sessionUserSubscription = this.session.sessionUser.subscribe(
            valor => {
                this.sessionUser = valor
                if (this.sessionLogin.isLoggedIn()) {
                    if (this.sessionUser.isEmpty()) {
                        this.downloadUser(true)
                    } else {
                        this.isStudent = valor.getRol() == User.RolEnum.Estudiante
                        this.isTeacher = valor.getRol() == User.RolEnum.Profesor
                        this.isAdmin = valor.getRol() == User.RolEnum.Administrador
                    }
                }
            }
        )

        this.messageSubscription = this.data.mensajeActual.subscribe(
            valor => this.mensaje = valor
        )
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