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
    readonly userid: String
    constructor(logged:boolean, userid?:String){
        this.logged = logged
        this.userid = userid || ""
    }
}