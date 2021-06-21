import { User, Team, Question, Test, Answer, QuestionAnswer, EmailConfiguration, CustomizedView, Prueba } from "@javgat/devtest-api";

export enum Tipo {
    SUCCESS = "success",
    ERROR = "danger",
    WARNING = "warning",
    DOWNLOADING = "downloading",
    SENDING = "sending",
    INFO = "info"
}

export class Mensaje {
    readonly texto: String;
    readonly type: Tipo;
    readonly trueType: Tipo;
    readonly mostrar: boolean;
    constructor(mensaje?: String, type?: Tipo, mostrar?: boolean) {
        this.texto = mensaje || "";
        this.trueType = type || Tipo.INFO;
        this.type = type || Tipo.INFO;
        if(type==Tipo.SENDING || type==Tipo.DOWNLOADING)
            this.type = Tipo.WARNING
        this.mostrar = mostrar || false;
    }

    showSpinner(): boolean{
        return this.isLoading()
    }
    
    showCheckCircle(): boolean{
        return this.trueType == Tipo.SUCCESS
    }

    showExclamationTriangle(): boolean{
        return (this.trueType == Tipo.ERROR) || (this.trueType == Tipo.WARNING)
    }

    showInfo(): boolean{
        return this.trueType == Tipo.INFO
    }

    isLoading(): boolean{
        return this.trueType==Tipo.SENDING || this.trueType==Tipo.DOWNLOADING
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
    readonly tiporol: string;
    /**
     * 
     * @param username If undefined, SessionUser isEmpty
     * @param email If undefined, SessionUser isEmpty
     * @param fullname 
     * @param type 
     */
    constructor(username?: string, email?: string, fullname?: string, rol?: User.RolEnum, tiporol?: string) {
        this.username = username || ""
        this.email = email || ""
        this.fullname = fullname || ""
        this.rol = rol || User.RolEnum.Estudiante
        this.tiporol = tiporol || ""
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

    getTipoRol(): string{
        return this.tiporol
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
    constructor(username: string, email: string, fullname: string, rol: User.RolEnum, tiporol: string | undefined) {
        super(username, email, fullname, rol, tiporol)
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
    id: number;
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
    isRespondida: boolean;
    penalizacion: number;
    cantidadFavoritos: number;
    posicion: number;
    prevId: number;
    nextId: number;

    constructor(id?: number | undefined, title?: string, question?: string, estT?: number, autoC?: boolean,
        edit?: boolean, username?: string, accesoPublicoNoPublicada?: boolean, eleUni?: boolean | undefined, solu?: string | undefined,
        tipo?: Question.TipoPreguntaEnum, valor?: number | undefined, isRespondida?: boolean, penalizacion?: number,
        cantidadFavoritos?: number, posicion?: number, prevId?: number, nextId?: number) {
        this.id = id || 0
        this.title = title || ""
        this.question = question || ""
        this.estimatedTime = estT || 0
        this.autoCorrect = autoC || false
        this.editable = edit || false
        this.username = username || ""
        this.eleccionUnica = eleUni
        this.solucion = solu
        this.tipoPregunta = tipo || Question.TipoPreguntaEnum.String
        this.valorFinal = valor
        this.accesoPublicoNoPublicada = accesoPublicoNoPublicada || false
        this.isRespondida = isRespondida || false
        this.penalizacion = penalizacion || 0
        this.cantidadFavoritos = cantidadFavoritos || 0
        this.posicion = posicion || 0
        this.prevId = prevId || 0
        this.nextId = nextId || 0
    }

    static constructorFromQuestion(q: Question): Pregunta{
        return new Pregunta(q.id, q.title, q.question, q.estimatedTime, q.autoCorrect, q.editable, q.username,
        q.accesoPublicoNoPublicada, q.eleccionUnica, q.solucion, q.tipoPregunta, q.valorFinal, q.isRespondida,
        q.penalizacion, q.cantidadFavoritos, q.posicion, q.prevId, q.nextId)
    }

    hasPrevious(): boolean{
        return this.prevId!=-1
    }

    hasNext(): boolean{
        return this.nextId!=-1
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
    horaCreacion?: Date;
    cantidadRespuestasDelUsuario: number;
    notaMaxima: number;
    autoCorrect: boolean;
    visibilidad: Test.VisibilidadEnum;
    cantidadFavoritos: number;
    tiempoEstricto: boolean;
    maxIntentos: number;
    constructor(title?: string, description?: string, accesoPublico?: boolean, editable?: boolean, maxMinutes?: number, 
        username?: string, id?: number, accesoPublicoNoPublicado?: boolean, horaCreacion?: Date, originalTestID?: number,
        cantidadRespuestasDelUsuario?: number, notaMaxima?: number, autoCorrect?: boolean, visibilidad?: Test.VisibilidadEnum,
        cantidadFavoritos?: number, tiempoEstricto?: boolean, maxIntentos?: number){
        this.title=title || ""
        this.description=description || ""
        this.maxMinutes=maxMinutes || 0
        this.accesoPublico=accesoPublico || false
        this.editable=editable || false
        this.username=username || ""
        this.id=id || 0
        this.accesoPublicoNoPublicado=accesoPublicoNoPublicado || false
        this.originalTestID = originalTestID || -1
        this.horaCreacion = horaCreacion
        this.cantidadRespuestasDelUsuario = cantidadRespuestasDelUsuario || 0
        this.notaMaxima = notaMaxima || 0
        this.autoCorrect = autoCorrect || false
        this.visibilidad = visibilidad || Test.VisibilidadEnum.Manual
        this.cantidadFavoritos = cantidadFavoritos || 0
        this.tiempoEstricto = tiempoEstricto || false
        this.maxIntentos = maxIntentos || 0
    }

    static constructorFromTest(t: Test): Examen{
        return new Examen(t.title, t.description, t.accesoPublico, t.editable, t.maxMinutes, t.username, t.id,
            t.accesoPublicoNoPublicado, t.horaCreacion, t.originalTestID, t.cantidadRespuestasDelUsuario, t.notaMaxima,
            t.autoCorrect, t.visibilidad, t.cantidadFavoritos, t.tiempoEstricto, t.maxIntentos)
    }

    static visibilidadToString(vis: Test.VisibilidadEnum ) : string{
        switch(vis){
            case Test.VisibilidadEnum.Manual:
                return "Manual"
            case Test.VisibilidadEnum.AlEntregar:
                return "Al entregar"
            case Test.VisibilidadEnum.AlCorregir:
                return "Al corregir"
            default:
                return ""
        }
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
                return "Elección múltiple con respuesta única"
            else
                return "Elección múltiple con respuesta múltiple"
        default:
            return ""
    }
}

export function bgcolorQAnswerPuntuacion(puntuacion: number): string{
    if(puntuacion >= 100){
        return "bg-success"
    }else if(puntuacion >0){
        return "bg-warning text-dark"
    }else{
        return "bg-danger"
    }
}

export class Respuesta implements Answer{
    id: number;
    startTime?: Date | undefined;
    finishTime?: Date | undefined;
    entregado: boolean;
    testid: number;
    username: string;
    puntuacion: number;
    corregida: boolean;
    visibleParaUsuario: boolean;
    constructor(resp?: Answer){
        this.id = resp?.id || 0
        this.startTime = resp?.startTime
        this.finishTime = resp?.finishTime
        this.entregado = resp?.entregado || false
        this.testid = resp?.testid || 0
        this.username = resp?.username || ""
        this.puntuacion = resp?.puntuacion || 0
        this.corregida = resp?.corregida || false
        this.visibleParaUsuario = resp?.visibleParaUsuario || false
    }
}

export class RespuestaPregunta implements QuestionAnswer{
    respuesta: string;
    idPregunta: number;
    idRespuesta: number;
    puntuacion: number;
    corregida: boolean;
    indicesOpciones: number[];
    username: string;
    estado: QuestionAnswer.EstadoEnum;
    constructor(qa?: QuestionAnswer){
        this.idPregunta = qa?.idPregunta || 0
        this.idRespuesta = qa?.idRespuesta || 0
        this.puntuacion = qa?.puntuacion || 0
        this.corregida = qa?.corregida || false
        this.respuesta = qa?.respuesta || ""
        this.indicesOpciones = qa?.indicesOpciones || []
        this.username = qa?.username || ""
        this.estado = qa?.estado || QuestionAnswer.EstadoEnum.NoProbado
    }

}

export class ConfiguracionCorreo implements EmailConfiguration{
    from: string;
    password: string;
    serverhost: string;
    serverport: number;
    frontendurl: string;
    constructor(c?: ConfiguracionCorreo){
        this.from = c?.from || ""
        this.password = c?.password || ""
        this.serverhost = c?.serverhost || ""
        this.serverport = c?.serverport || 0
        this.frontendurl = c?.frontendurl || ""
    }
}

export enum EnumOrderBy {
    newDate = "newDate",
    oldDate = "oldDate",
    moreFav = "moreFav",
    lessFav = "lessFav",
    moreTime = "moreTime",
    lessTime = "lessTime"
}

export enum EnumOrderByAnswer {
    newStartDate = "newStartDate",
    oldStartDate = "oldStartDate",
    morePuntuacion = "morePuntuacion",
    lessPuntuacion = "lessPuntuacion",
    moreDuracion = "moreDuracion",
    lessDuracion = "lessDuracion",
}

export class VistaPersonalizada implements CustomizedView{
    rolBase: CustomizedView.RolBaseEnum;
    mensajeInicio: string;
    constructor(c?: CustomizedView){
        this.rolBase = c?.rolBase || CustomizedView.RolBaseEnum.NoRegistrado
        this.mensajeInicio = c?.mensajeInicio || ""
    }
}

export class PruebaEjecucion implements Prueba{
    id?: number | undefined;
    preguntaid?: number | undefined;
    entrada: string;
    salida: string;
    visible: boolean;
    postEntrega: boolean;
    valor: number;
    constructor(p?: Prueba){
        this.id = p?.id || 0
        this.preguntaid = p?.preguntaid || 0
        this.entrada = p?.entrada || ""
        this.salida = p?.salida || ""
        this.visible = p?.visible || false
        this.postEntrega = p?.postEntrega || false
        this.valor = p?.valor || 0
    }
}