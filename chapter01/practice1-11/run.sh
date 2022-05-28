#!/bin/bash

readonly URLS=(
"http://code-examples.net/"
"http://kotaeta.com/"
"http://answer-id.com/"
"http://de-vraag.com/"
"http://code.i-harness.com/"
"http://living-sun.com/"
"http://qastack.jp/"
"http://qastack.ru/"
"http://qastack.it/"
"http://qastack.mx/"
"http://qastack.com.br/"
"http://qastack.info.tr/"
"http://qastack.in.th/"
"http://qastack.com.de/"
"http://qastack.fr/"
"http://qastack.cn/"
"http://qastack.com.ua/"
"http://qastack.co.in/"
"http://qastack.kr/"
"http://qastack.vn/"
"http://qastack.net.bd/"
"http://qa-stack.pl/"
"http://qastack.id/"
"http://www.coder.work/"
"http://www.it-swarm-ja.tech/"
"http://www.it-swarm.jp.net/"
"http://www.it-swarm-ja.com/"
"http://www.webdevqa.jp.net/"
"http://www.web-dev-qa-db-ja.com/"
"http://www.it-swarm-fr.com/"
"http://www.web-dev-qa-db-fr.com/"
"http://codeflow.site/"
"http://codeguides.site/"
"http://overcoder.net/"
"http://coderoad.ru/"
"http://www.javaer101.com/"
"http://voidcc.com/"
"http://siwib.org/"
"http://fluffyfables.com/"
"http://www.fixes.pub/"
"http://knews.vip/"
"http://isolution.pro/"
"http://uwenku.com/"
"http://athabasca-foto.com/"
"http://zsharp.org/"
"http://projectbackpack.org/"
"http://waymanamechurch.org/"
"http://sunflowercreations.org/"
"http://cfadnc.org/"
"http://fitforlearning.org/"
"http://panaindustrial.com/"
"http://sierrasummit2005.org/"
"http://theshuggahpies.com/"
"http://pcbconline.org/"
"http://www.nuomiphp.com/"
"http://ubuntu.buildwebhost.com/"
"http://ubuntuaa.com/"
"http://www.debugcn.com/"
"http://sch22.org/"
"http://gupgallery.com/"
"http://amuddycup.com/"
"http://ecnf2016.org/"
"http://softwareuser.asklobster.com/"
"http://domainelespailles.net/"
"http://ec-europe.org/"
"http://pakostnici.com/"
"http://try2explore.com/"
"http://itectec.com/"
"http://stackovergo.com/"
"http://faithcov.org/"
"http://noblenaz.org/"
"http://culinarydegree.info/"
"http://qapicks.com/"
"http://narkive.jp/"
"http://ourladylakes.org/"
"http://intellipaat.com/"
"http://newbedev.com/"
"http://www.codenong.com/"
"http://routinepanic.com/"
"http://tousu.in/"
"http://tutorialmore.com/"
"http://www.titanwolf.org/"
"http://coderedirect.com/"
"http://fullstackuser.com/"
"http://ostack.cn/"
"http://webdevdesigner.com/"
"http://www.ghcc.net/"
"http://developreference.com/"
"http://www.semicolonworld.com/"
"http://tipsfordev.com/"
"http://www.qi-u.com/"
"http://www.xsprogram.com/"
"http://stackoom.com/"
"http://cndgn.com/"
"http://www.generacodice.com/"
"http://stackfinder.jp.net/"
"http://www.uebu-kaihatsu.jp.net/"
"http://sqlite.in/"
"http://stackguides.com/"
"http://younggeeks.in/"
"http://www.answerlib.com/"
"http://edupro.id/"
"http://www.stackfinder.ru/"
"http://www.desenv-web-rp.com/"
"http://www.web-dev-qa-db-pt.com/"
"http://www.jscodetips.com/"
"http://www.5axxw.com/"
"http://question-it.com/"
"http://codefaq.info/"
"http://codefaq.ru/"
"http://mediatagtw.com/"
"http://progi.pro/"
"http://jike.in/"
"http://elfishgene.com/"
"http://sysadminde.com/"
"http://answacode.com/"
"http://ask-dev.ru/"
"http://arip-photo.org/"
"http://jablogs.com/"
"http://jpdebug.com/"
"http://askcodez.com/"
"http://iquestion.pro/"
"http://ntcdoon.org/"
"http://programmierfrage.com/"
"http://microeducate.tech/"
"http://www.debugko.com/"
"http://devdreamz.com/"
"http://catwolf.org/"
"http://1r1g.com/"
"http://string.quest/"
"http://reddit.fun/"
"http://qa.icopy.site/"
"http://errorsfixing.com/"
"http://syntaxfix.com/"
"http://codegrepr.com/"
"http://quabr.com/"
"http://serveanswer.com/"
"http://safehavenpetrescue.org/"
"http://cainiaojiaocheng.com/"
"http://www.it-mure.jp.net/"
"http://www.it-swarm.com.ru/"
"http://www.it-roy-ru.com/"
"http://brocante.dev/"
"http://laravelquestions.com/"
"http://www.py4u.net/"
"http://issues-world.com/"
"http://article.docway.net/"
"http://codehero.jp/"
"http://pretagteam.com/"
"http://lycaeum.dev/"
"http://classmethod.dev/"
"http://stormcrow.dev/"
"http://christfever.in/"
"http://coredump.biz/"
"http://stackify.dev/"
"http://answerlib.com/"
"http://fooobar.com/"
"http://askdev.vn/"
"http://devbugfix.com/"
"http://www.generacodice.blog/"
"http://www.generacodice.it/"
"http://coder-solution-jp.com/"
"http://jpndev.com/"
"http://stackfault.net/"
"http://alltodev.com/"
"http://www.buzzphp.com/"
"http://quares.ru/"
"http://ibootweb.com/"
"http://howtofix.io/"
"http://www.examplefiles.net/"
"http://www.web-dev-qa.com/"
"http://exceptionshub.com/"
"http://easysavecode.com/"
"http://python-stack.de/"
"http://xiu2.net/"
"http://pythonwd.com/"
"http://stackqna.com/"
"http://codersatellite.com/"
"http://codewdw.com/"
"http://jpcodeqa.com/"
)

go run $(pwd)/practice1-11/main.go ${URLS[@]}