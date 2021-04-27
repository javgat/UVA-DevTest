import { User, Team, Question, Test } from "@javgat/devtest-api";

export enum Tipo {
    SUCCESS = "success",
    ERROR = "danger",
    WARNING = "warning",
    INFO = "info"
}

export class Mensaje {
    readonly texto: String;
    readonly type: Tipo;
    readonly mostrar: boolean;
    constructor(mensaje?: String, type?: Tipo, mostrar?: boolean) {
        this.texto = mensaje || "";
        this.type = type || Tipo.INFO;
        this.mostrar = mostrar || false;
    }
}

export class SessionLogin {
    private readonly logged: boolean
    private readonly username: string
    constructor(logged: boolean, username?: string) {
        this.logged = logged
        this.username = username || ""
    }

    isLoggedIn(): boolean {
        return this.logged
    }

    getUserUsername(): string {
        return this.username
    }
}

export class SessionUser implements User {
    readonly empty: boolean
    readonly username: string;
    readonly email: string;
    readonly fullname: string;
    readonly rol: User.RolEnum;
    /**
     * 
     * @param username If undefined, SessionUser isEmpty
     * @param email If undefined, SessionUser isEmpty
     * @param fullname 
     * @param type 
     */
    constructor(username?: string, email?: string, fullname?: string, rol?: User.RolEnum) {
        this.username = username || ""
        this.email = email || ""
        this.fullname = fullname || ""
        this.rol = rol || User.RolEnum.Estudiante
        this.empty = (this.username == "" || this.email == "")
    }

    isEmpty(): boolean {
        return this.empty
    }

    getUsername(): string {
        return this.username
    }

    getEmail(): string {
        return this.email
    }

    getFullname(): string {
        return this.fullname
    }

    getRol(): User.RolEnum {
        return this.rol
    }

    isStudent(): boolean {
        return this.getRol() == User.RolEnum.Estudiante
    }

    isTeacher(): boolean {
        return this.getRol() == User.RolEnum.Profesor
    }

    isAdmin(): boolean {
        return this.getRol() == User.RolEnum.Administrador
    }

    isTeacherOrAdmin(): boolean {
        return (this.isTeacher() || this.isAdmin())
    }
}

export class Usuario extends SessionUser {
    /**
     * 
     * @param username
     * @param email
     * @param fullname 
     * @param type 
     */
    constructor(username: string, email: string, fullname: string, rol: User.RolEnum) {
        super(username, email, fullname, rol)
    }
}

export class Equipo implements Team {
    teamname: string;
    description: string;
    soloProfesores: boolean;
    constructor(teamname: string, description: string, soloProfesores: boolean) {
        this.teamname = teamname
        this.description = description
        this.soloProfesores = soloProfesores
    }

    getTeamname(): string {
        return this.teamname
    }

    getDescription(): string {
        return this.description
    }

    isSoloProfesores(): boolean {
        return this.soloProfesores
    }
}

export class Pregunta implements Question {
    id: number | undefined;
    title: string;
    question: string;
    estimatedTime: number;
    autoCorrect: boolean;
    editable: boolean;
    username: string;
    eleccionUnica: boolean | undefined;
    solucion: string | undefined;
    tipoPregunta: Question.TipoPreguntaEnum;
    valorFinal: number | undefined;
    accesoPublicoNoPublicada: boolean;

    constructor(id: number | undefined, title: string, question: string, estT: number, autoC: boolean,
        edit: boolean, username: string, accesoPublicoNoPublicada: boolean, eleUni: boolean | undefined, solu: string | undefined,
        tipo: Question.TipoPreguntaEnum, valor: number | undefined) {
        this.id = id
        this.title = title
        this.question = question
        this.estimatedTime = estT
        this.autoCorrect = autoC
        this.editable = edit
        this.username = username
        this.eleccionUnica = eleUni
        this.solucion = solu
        this.tipoPregunta = tipo
        this.valorFinal = valor
        this.accesoPublicoNoPublicada = accesoPublicoNoPublicada
    }
}

export class Examen implements Test{
    id: number;
    title: string;
    description: string;
    maxMinutes: number;
    accesoPublico: boolean;
    editable: boolean;
    username: string;
    accesoPublicoNoPublicado: boolean;
    originalTestID: number;
    horaCreacion: string;
    cantidadRespuestasDelUsuario: number;
    constructor(title?: string, description?: string, accesoPublico?: boolean, editable?: boolean, maxMinutes?: number, 
        username?: string, id?: number, accesoPublicoNoPublicado?: boolean, horaCreacion?: string, originalTestID?: number,
        cantidadRespuestasDelUsuario?: number){
        this.title=title || ""
        this.description=description || ""
        this.maxMinutes=maxMinutes || 0
        this.accesoPublico=accesoPublico || false
        this.editable=editable || false
        this.username=username || ""
        this.id=id || 0
        this.accesoPublicoNoPublicado=accesoPublicoNoPublicado || false
        this.originalTestID = originalTestID || -1
        this.horaCreacion = horaCreacion || ""
        this.cantidadRespuestasDelUsuario = cantidadRespuestasDelUsuario || 0
    }

    static constructorFromTest(t: Test): Examen{
        return new Examen(t.title, t.description, t.accesoPublico, t.editable, t.maxMinutes, t.username, t.id,
            t.accesoPublicoNoPublicado, t.horaCreacion, t.originalTestID, t.cantidadRespuestasDelUsuario)
    }
}

export function tipoPrint(tipo: string, eleccionUnica: boolean | undefined): string {
    switch (tipo) {
        case Question.TipoPreguntaEnum.String:
            return "Texto"
        case Question.TipoPreguntaEnum.Codigo:
            return "Código"
        case Question.TipoPreguntaEnum.Opciones:
            if (eleccionUnica)
                return "Respuesta única"
            else
                return "Respuesta múltiple"
        default:
            return ""
    }
}