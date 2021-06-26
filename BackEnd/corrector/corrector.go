// UVa-DevTest. 2021.
// Author: Javier Gatón Herguedas.

// Package corrector executes Pruebas for the Code QuestionAnswers
package corrector

import (
	"errors"
	"log"
	"strconv"
	"strings"
	"uva-devtest/models"
	"uva-devtest/persistence/dao"
	"uva-devtest/persistence/dbconnection"
)

const main_directory = "/tmp/pruebas/"

func executePruebas(answerid int64, questionid int64, isPostEntrega bool) {
	var sumaValores int64 = 0
	var valoresObtenidos int64 = 0
	db, err := dbconnection.ConnectDb()
	errorInterno := "Error interno del servidor"
	if err == nil {
		sharedDir := main_directory + strconv.FormatInt(answerid, 10) + "_" + strconv.FormatInt(questionid, 10) + "/"
		err = prepareFolder(sharedDir)
		if err == nil {
			var ps []*dao.Prueba
			ps, err = dao.GetPruebas(db, questionid)
			if err == nil {
				for _, p := range ps {
					if !*p.PostEntrega || isPostEntrega {
						err = writePruebaInput(p, sharedDir)
					}
					if err != nil {
						break
					}
				}
				if err == nil {
					var qa *dao.QuestionAnswer
					qa, err = dao.GetQuestionAnswerFromAnswer(db, answerid, questionid)
					if err == nil {
						err = writeCodeProgram(qa.Respuesta, sharedDir)
						if err == nil {
							err = dao.DeleteEjecuciones(db, answerid, questionid)
							if err == nil {
								err = execDocker(sharedDir)
								if err == nil {
									var errsComp string
									errsComp, err = readCompilationErrors(sharedDir)
									if err == nil {
										var puntuacion int64 = 0
										if errsComp == "" {
											for _, p := range ps {
												if !*p.PostEntrega || isPostEntrega {
													var ej *dao.Ejecucion
													ej, err = readPruebaInput(p, answerid, questionid, sharedDir)
													if err == nil {
														var est string
														sumaValores += *p.Valor
														if !strings.EqualFold(strings.Trim(*ej.Estado, " "), strings.Trim(models.PruebaEstadoErrorRuntime, " ")) &&
															!strings.EqualFold(strings.Trim(*ej.Estado, " "), strings.Trim(models.PruebaEstadoTiempoExcedido, " ")) {
															if *ej.SalidaReal == *p.Salida {
																est = dao.EstadoEjecucionCorrecto
																valoresObtenidos += *p.Valor
															} else {
																est = dao.EstadoEjecucionSalidaIncorrecta
															}
															ej.Estado = &est
														}
														err = dao.SaveEjecucion(db, ej)
													}
												}
												if err != nil {
													break
												}
											}
											err = dao.SetQuestionAnswerProbado(db, answerid, questionid)
											if sumaValores > 0 {
												puntuacion = (valoresObtenidos * 100) / sumaValores
											} else {
												puntuacion = 0
											}
										} else {
											err = dao.SetQuestionAnswerErrorCompilacion(db, &errsComp, answerid, questionid)
											puntuacion = 0
										}
										if err == nil && isPostEntrega {
											review := &models.Review{
												Puntuacion: &puntuacion,
											}
											err = UpdateReview(answerid, questionid, review)
										}
									}
								}
							}

						}
					}
				}
			}
		}
		if err != nil {
			log.Println("Error en corrector en ExecutePruebas(): ", err)
			err = dao.SetQuestionAnswerErrorCompilacion(db, &errorInterno, answerid, questionid)
			if err != nil {
				log.Println("Error en corrector en ExecutePruebas(): ", err)
			}
		}
		err = cleanFolder(sharedDir)
	}

	if err != nil {
		log.Println("Error en corrector en ExecutePruebas(): ", err)
		err = dao.SetQuestionAnswerErrorCompilacion(db, &errorInterno, answerid, questionid)
		if err != nil {
			log.Println("Error en corrector en ExecutePruebas(): ", err)
		}
	}
}

func ExecutePrePruebas(answerid int64, questionid int64) {
	executePruebas(answerid, questionid, false)
}

func ExecuteFullPruebas(answerid int64, questionid int64) {
	executePruebas(answerid, questionid, true)
}

func AddAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var a *dao.Answer
		var q *dao.Question
		a, err = dao.GetAnswer(db, aid)
		if err == nil {
			if a == nil {
				return errors.New("no se encontro el recurso")
			}
			q, err = dao.GetQuestionFromTest(db, a.Testid, qid)
			if err == nil {
				if q == nil {
					return errors.New("no se encontro el recurso")
				}
				punt := a.Puntuacion + float64(*q.ValorFinal*puntuacion)/float64(100)
				err = dao.PutAnswerPuntuacion(db, aid, punt)
				if err == nil {
					return nil
				}
			}
		}
	}
	return err
}

func SubstractAnswerPuntuacion(aid int64, qid int64, puntuacion int64) error {
	return AddAnswerPuntuacion(aid, qid, -puntuacion)
}

func UpdateReview(aid int64, qid int64, review *models.Review) error {
	db, err := dbconnection.ConnectDb()
	if err == nil {
		var qa *dao.QuestionAnswer
		qa, err = dao.GetQuestionAnswerFromAnswer(db, aid, qid)
		if qa != nil && err == nil {
			err = SubstractAnswerPuntuacion(aid, qid, *qa.Puntuacion)
		}
		if err == nil {
			log.Println(*review.Puntuacion)
			err = AddAnswerPuntuacion(aid, qid, *review.Puntuacion)
			if err == nil {
				err = dao.PutReview(db, aid, qid, review)
			}
		}
	}
	return err
}
