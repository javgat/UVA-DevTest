export enum Tipo{
    SUCCESS = "success",
    ERROR = "danger",
    WARNING= "warning",
    INFO="info"
}

export class Mensaje{
    readonly texto: String;
    readonly type : Tipo;
    readonly mostrar : boolean;
    constructor(mensaje?:String, type?:Tipo, mostrar?:boolean){
        this.texto = mensaje || "";
        this.type = type || Tipo.INFO;
        this.mostrar = mostrar || false;
    }
}

export class SessionUser{
    readonly logged: boolean
    readonly jwt: String
    readonly userid: String
    constructor(logged:boolean, jwt?:String, userid?:String){
        this.logged = logged
        this.jwt = jwt || ""
        this.userid = userid || ""
    }
}