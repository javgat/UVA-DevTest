/**
 * DevTest
 * DevTest
 *
 * OpenAPI spec version: 1.0.0
 * Contact: javigaton@gmail.com
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 * Do not edit the class manually.
 */


export interface QuestionAnswer { 
    respuesta?: string;
    idPregunta: number;
    idRespuesta: number;
    puntuacion: number;
    corregida: boolean;
    indicesOpciones?: Array<number>;
}