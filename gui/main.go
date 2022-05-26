package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/pavelerokhin/go-and-scrape/gui/cli"
	"github.com/pavelerokhin/go-and-scrape/models/entities"
)

func main() {
	logger := log.New(os.Stdout, "gui ", log.LstdFlags|log.Lshortfile)

	var mode string
	if len(os.Args) < 2 {
		//logger.Fatalln("no GUI mode has been set")
		mode = "cli"
	}
	//mode = os.Args[1]

	if mode == "cli" {

		news := []entities.ArticlePreview{
			{
				Title:    "Российское юрлицо Google собирается объявить о банкротстве",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "10:38, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Компания ООО «Гугл» намерена обратиться в суд с заявлением о признании себя банкротом,  следует из опубликованного ( https://fedresurs.ru/sfactmessage/B67464A6A16845AB909F2B5122CE6AFE ) ею уведомления.  На это обратил внимание «Интерфакс».  «С 22 марта 2022 года [компания] предвидит собственное банкротство и невозможность исполнения денежных обязательств»,  — сказано в документе.  В самой компании происходящее пока не прокомментировали.  Как пишет «Интерфакс»,  по итогам 2021 года ООО «Гугл» впервые за долгое время получило чистый убыток.  Он составил 26 миллиардов рублей.  В декабре 2021 года суд назначил компании «Гугл» оборотный (то есть зависимый от выручки) штраф в размере 7, 2 миллиарда рублей за неудаление запрещенного контента.  В мае приставы завели дело о принудительном взыскании штрафа.  В 2022 году суд дважды арестовал имущество ООО «Гугл» на 500 миллионов рублей в качестве обеспечительных мер по искам о блокировке ютьюб-каналов НТВ и ТНТ.  «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Леонид Слуцкий возглавил фракцию ЛДПР в Госдуме",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "07:27, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Депутаты Госдумы от ЛДПР избрали новым руководителем фракции Леонида Слуцкого,  сообщает «Интерфакс» ( https://www.interfax.ru ).  Слуцкий — председатель комитета по международным делам Госдумы — временно исполнял обязанности главы фракции ЛДПР после смерти лидера партии Владимира Жириновского ( https://meduza.io/feature/2022/04/06/vse-chto-pridumal-zhirinovskiy-delaet-seychas-vlast ) ,  скончавшегося 4 апреля от последствий коронавирусной инфекции.  Леонида Слуцкого обвиняли в домогательствах ------------------------------------------- * Друг патриарха,  присоединитель Крыма Как депутат Леонид Слуцкий,  пристававший к журналисткам,  добился успеха и почему государство его защищает ( /feature/2018/03/23/drug-patriarha-prisoedinitel-kryma ) * Я никого не называл «зайчуткой».  Я обычно говорю «зайчонок» Первое большое интервью депутата Леонида Слуцкого после обвинений в домогательствах.  Максимально коротко ( /paragraph/2018/06/13/ya-nikogo-ne-nazyval-zaychutkoy-ya-obychno-govoryu-zaychonok ) * Журналистки обвинили Леонида Слуцкого в домогательствах.  Они могут подать в суд? А что с депутатской неприкосновенностью? ( /feature/2018/02/28/zhurnalistki-obvinili-leonila-slutskogo-v-domogatelstvah-oni-mogut-podat-v-sud-a-chto-s-deputatskoy-neprikosnovennostyu ) 27 мая состоится съезд ЛДПР,  на котором будет избран новый глава партии.  Подробнее о претендентах на пост главы ЛДПР В Кремле выбирают преемника Владимира Жириновского.  Среди кандидатов — депутат Леонид Слуцкий и телеведущий Владимир Соловьев Правда,  в администрации президента уже думают о «постепенной утилизации» ЛДПР ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- месяц назад ( /feature/2022/04/08/v-kremle-vybirayut-preemnika-vladimira-zhirinovskogo-sredi-kandidatov-deputat-leonid-slutskiy-i-televeduschiy-vladimir-soloviev ) Подробнее о претендентах на пост главы ЛДПР В Кремле выбирают преемника Владимира Жириновского.  Среди кандидатов — депутат Леонид Слуцкий и телеведущий Владимир Соловьев Правда,  в администрации президента уже думают о «постепенной утилизации» ЛДПР ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- месяц назад ( /feature/2022/04/08/v-kremle-vybirayut-preemnika-vladimira-zhirinovskogo-sredi-kandidatov-deputat-leonid-slutskiy-i-televeduschiy-vladimir-soloviev ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Александр Роднянский снимет сериал про Путина.  Он будет основан на книге «Вся кремлевская рать» Михаила Зыгаря",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "09:44, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Продюсер Александр Роднянский объявил о работе над сериалом про российского президента Владимира Путина,  который будет основан на книге Михаила Зыгаря «Вся кремлевская рать».  Как пишет Variety, сериал будет посвящен тому,  как «невзрачный бывший офицер КГБ превратился в одного из самых устрашающих политиков в мире».  По словам Роднянского,  сравнившего сериал с «Карточным домиком»,  в центре повествования будет внутренняя эволюция Путина.  «„Вся кремлевская рать“ покажет международной аудитории шаг за шагом,  как разум Путина растворился в теориях заговора,  и как люди вокруг него использовали его глубокие страхи и недоверие западным ценностям,  которые он всегда считал циничной ложью»,  — заявил продюсер.  Информации о возможной дате выхода сериала пока нет.  Продюсер Александр Роднянский,  уроженец Киева и один из самых влиятельных деятелей российского кино последних десятилетий,  с началом российского вторжения в Украину уехал из России и открыто высказывается против войны.  Читайте также ------------- * «Радио Долин»: как патриотические блокбастеры годами готовили зрителей к войне и почему кино все же не стало главным рупором пропаганды Интервью Александра Роднянского ( /feature/2022/04/13/radio-dolin-kak-patrioticheskie-blokbastery-godami-gotovili-zriteley-k-voyne-i-pochemu-kino-vse-zhe-ne-stalo-glavnym-ruporom-propagandy ) * Продюсерская компания Александра Роднянского будет снимать сериалы для Apple TV+ ( /news/2021/07/12/prodyuserskaya-kompaniya-aleksandra-rodnyanskogo-budet-snimat-serialy-dlya-apple-tv ) * Режиссер «Наследников» Андрий Парекх снимет сериал продюсера Александра Роднянского «Красная радуга» о гей-активистах в СССР ( /news/2021/12/08/rezhisser-naslednikov-andriy-parekh-snimet-serial-prodyusera-aleksandra-rodnyanskogo-krasnaya-raduga-o-gey-aktivistah-v-sssr ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Лукашенко подписал закон о смертной казни за покушение на совершение теракта.  По этой статье в Беларуси обвиняют оппозиционеров",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "09:03, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Александр Лукашенко подписал закон,  позволяющий приговаривать к смертной казни за покушение на совершение терактов.  Документ опубликован на национальном правовом интернет-портале.  Поправки,  сообщает «Интерфакс» ( https://www.interfax.ru/ ) ,  вступят в силу через 10 дней после официального опубликования.  В качестве исключительной меры наказания допускается применение смертной казни — расстрела (до отмены смертной казни) за преступления ,  предусмотренные частью 2 статьи 124,  частью 3 статьи 126,  частью 3 статьи 289 и частью 2 статьи 359 настоящего Кодекса,  либо за сопряженные с умышленным лишением жизни человека при отягчающих обстоятельствах,  некоторые иные особо тяжкие преступления.  По статье 289 «Акт терроризма» в Беларуси возбуждены уголовные дела в отношении некоторых представителей оппозиции,  включая Светлану Тихановскую — соперницу Лукашенко на последних на настоящий момент президентских выборах.  Ранее в Беларуси нельзя было приговаривать к смертной казни за приготовление преступлений.  В конце апреля соответствующие поправки одобрила нижняя палата парламента Беларуси,  а в начале мая — Совет республики.  Беларусь — единственная страна Европы,  применяющая смертную казнь.  Подробнее об этом Отбыл по приговору Белоруссия — единственная страна Европы,  где есть смертная казнь.  Саша Сулим рассказывает,  как государство убивает преступников -------------------------------------------------------------------------------------------------------------------------------------------------- 4 года назад ( /feature/2018/02/02/otbyl-po-prigovoru ) Подробнее об этом Отбыл по приговору Белоруссия — единственная страна Европы,  где есть смертная казнь.  Саша Сулим рассказывает,  как государство убивает преступников -------------------------------------------------------------------------------------------------------------------------------------------------- 4 года назад ( /feature/2018/02/02/otbyl-po-prigovoru ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "ФСБ отчиталась о задержании «сторонника украинских нацистов»,  подозреваемого в повреждении ЛЭП в Кемерово",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "07:01, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) В Кемерово задержан местный житель,  подозреваемый в диверсии на линии электропередач,  сообщает Федеральная служба безопасности РФ.  В заявлении спецслужбы,  которое цитирует «Интерфакс»,  задержанный россиянин называется «сторонником украинских нацистов».  Утверждается,  что он причастен к повреждению двух опор ЛЭП,  в результате которого произошло временное отключение электроэнергии в Кемеровской области.  Насколько серьезным был сбой,  не уточняется.  В отношении задержанного,  чье имя не называется,  возбуждено уголовное дело по статьям «умышленное повреждение имущества» (часть 2 статьи 167 УК,  наказывается ( http://www.consultant.ru/document/cons_doc_LAW_10699/d260e55e06d1e6bc720d2e591a8383a43b1a5eed/ ) лишением свободы на срок до пяти лет) и «диверсия» (статья 281 УК,  наказывается ( https://www.consultant.ru/document/cons_doc_LAW_10699/26760b317cab8b7c84154df406492379d8f6abf1/ ) лишением свободы на срок от 10 до 20 лет или пожизненно).  Во время обысков у подозреваемого изъяли «бутылки с зажигательной смесью,  канистры с бензином,  крепежные болты с поврежденных ЛЭП,  два пневматических пистолета,  переделанных под стрельбу боевыми патронами,  два боевых ножа,  средства связи и коммуникации,  содержащие схемы расположения военкоматов и отделений полиции города Кемерово,  а также личный дневник с рукописными записями,  подтверждающими совершение им диверсий».  Задержанный дал признательные показания,  утверждают в ФСБ.  Главные новости дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- 2 часа назад ( /live/2022/05/18/voyna ) Главные новости дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- 2 часа назад ( /live/2022/05/18/voyna ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Как читать ресурсы,  заблокированные российскими властями",
				Subtitle: "Простая инструкция «Медузы»",
				Article: entities.Article{
					Author: "",
					Date:   "10:45, 26 февраля 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) В России началось «замедление» фейсбука.  По опыту аналогичных мер,  принятых против твиттера,  ясно,  что это может реально осложнить пользование социальной сетью — особенно неудобно будет смотреть видео и фотографии.  С высокой вероятностью на этом власти не остановятся.  Мы предполагаем,  что уже в ближайшее время они могут начать блокировку независимых СМИ (и не только).  Рассказываем,  что можно сделать уже сейчас,  чтобы подготовиться к жизни в новой интернет-реальности.  > > Эта инструкция есть и в гугл-доке ( > https://docs.google.com/document/d/1X3MV7rPQp8IIdPKclGj_8O0WI1wK9_ET_FlOpw9oeT4/edit?usp=sharing > ) — чтобы ей удобнее было делиться со знакомыми.  А вот инструкция ( > https://docs.google.com/document/d/12WJkqzEgAKkuRmtv0nQK__oWOhEtkE_QcQL82baoSZg/edit?usp=sharing > ) (тоже в гугл-доке),  как обойти блокировку «Медузы» Научитесь обходить блокировку ----------------------------- Но сперва поищите мобильное приложение,  выпущенное заблокированным (или еще не заблокированным!) ресурсом Это самый простой способ.  Заблокировать мобильное приложение сложнее,  чем обычный сайт.  Особенно если разработчики заранее встроили в него инструменты обхода блокировок.  Но если его не получается найти или установить,  придется идти другим путем.  «Медуза» заблокирована в России.  Мы были к этому готовы — и продолжаем работать.  Несмотря ни на что --------------------------------------------------------------------------------------------------- Нам нужна ваша помощь как никогда.  Прямо сейчас.  Дальше всем нам будет еще труднее.  Мы независимое издание и работаем только в интересах читателей.  Хочу помочь ( https://mdza.io/yiNxPZ7Bs-k ) Хочу помочь Установите VPN Это универсальный способ обхода блокировок.  Он позволяет пропустить весь ваш трафик через зашифрованное соединение.  Единственное,  что могут сделать российские власти — попытаться заблокировать вашего VPN-провайдера.  Тогда придется искать новый или учиться поднимать VPN-сервис самостоятельно (да,  так тоже можно сделать ( https://getoutline.org/en-GB/ ) ).  Купите подписку на VPN, выбрав компанию,  которая обещает не делиться вашими данными и проходит независимый аудит.  Придется немного потрудиться: нужно найти соответствующую информацию на сайте VPN-провайдера.  Хорошая новость в том,  что искать не очень сложно,  потому что VPN-провайдеры охотно сообщают о таком аудите (ведь от этого зависит их репутация).  Проверяйте,  кто именно сделал аудит: это должна быть известная компания,  вроде Pricewaterhouse Coopers или Ernst & Young, информацию о которой тоже легко найти в интернете.  > > *Важно!* Мы не рекомендуем пользоваться бесплатным VPN-сервисами — они > могут зарабатывать на сборе и продаже ваших данных.  Хорошая новость: из этого правила есть исключение.  После блокировки «Медузы» ProtonVPN предложил свою помощь,  рассказав,  что уже много лет поддерживает граждан тех стран,  где государство систематически занимается интернет-цензурой: например,  Турции,  Ирана и Беларуси.  И поскольку заплатить за зарубежный сервис с российской карты может не получиться,  Proton предоставляет пользователям бесплатную ( https://go.getproton.me/aff_c?offer_id=26&aff_id=2898&url_id=426&aff_sub=instrc ) версию своего VPN. Миссия ProtonVPN аналогична целям ( https://meduza.io/pages/codex ) нашей работы — обеспечить свободу слова и свободное распространение информации.  Никакой рекламы,  никаких ограничений по объему трафика,  никаких лог-файлов,  по которым вас можно было бы вычислить.  Разработчики ProtonVPN допускают,  что их начнут блокировать в России,  и к этому тоже готовятся.  Установите расширение для браузера,  помогающее обходить блокировки (рекомендуем два,  сделанных специально для пользователей из России) VPN-расширение « Обход блокировок рунета ( https://antizapret.prostovpn.org/ ) » пытается обойти блокировку только для запрещенных сайтов из так называемого единого реестра запрещенной информации Российской Федерации.  Аналогичное расширение Censor Tracker ( https://censortracker.org/ ) от «Роскомсвободы» еще и предупреждает,  если посещаемый вами сайт считается в России организатором распространения информации (то есть обязан следить за своими пользователям и передавать информацию о них по запросу властей,  в том числе сообщения и весь расшифрованный трафик).  Попробуйте Tor-браузер Скачайте ( https://tor.eff.org/ ) и установите ( https://tor.calyxinstitute.org/ ) приложение на компьютер или андроид-телефон (владельцы айфонов и айпадов могут скачать браузер ( https://apps.apple.com/us/app/onion-browser/id519296448 ) от стороннего разработчика).  В конце 2021 года власти России начали бороться с этой анонимной распределенной сетью.  Чтобы обойти блокировку,  многие пользователи используют так называемые мосты — непубличные узлы сети Tor. Список из нескольких таких мостов можно получить прямо в настройках браузера,  на сайте проекта ( https://bridges.torproject.org/ ) или с помощью электронной почты (правда,  на письмо ответят,  только если вы отправили его с учетной записи Gmail или Riseup ( https://riseup.net/ ) ).  Вы получите несколько строк,  которые нужно будет вставить в соответствующее поле в настройках браузера — а после этого сможете подключиться к анонимной сети,  несмотря на блокировку.  Другие важные инструкции по пользованию интернетом в России ----------------------------------------------------------- * Как пользоваться телефоном и интернетом в современной России — и не подставиться.  Правила,  которые нужно соблюдать абсолютно всем ( /feature/2019/08/16/kak-polzovatsya-telefonom-i-internetom-v-sovremennoy-rossii-i-ne-podstavitsya-pravila-kotorye-nuzhno-soblyudat-absolyutno-vsem ) Найдите альтернативные каналы получения информации от СМИ и НКО --------------------------------------------------------------- Практически все СМИ и НКО работают на разных платформах.  Попробуйте найти (желательно заранее) и подписаться на заблокированный ресурс в телеграме,  инстаграме,  ютьюбе,  ВК,  твиттере и фейсбуке.  Подпишитесь,  даже если эти платформы кажутся вам не самыми удобными.  Там вас наверняка своевременно оповестят,  каким образом можно получить доступ к заблокированному ресурсу (если дойдет до такого).  И конечно,  старая добрая электронная почта.  Власти пока еще не блокировали новостные рассылки — даже для тех,  кто пользуется российскими почтовыми сервисами (то есть для пользователей,  у которых имейл зарегистрирован,  скажем,  на @yandex.ru или @mail.ru).  Обязательно подпишитесь на аккаунты «Медузы».  Вот три самые важные платформы: * Телеграм ( https://t.me/meduzalive ) * Инстаграм ( https://www.instagram.com/meduzapro/ ) * «Вечерняя Медуза» ( https://meduza.io/specials/daily ) > > И не забывайте,  что все,  кто поддерживают ( https://support.meduza.io/ ) «Медузу»,  > могут также получать рассылку Kit — о том,  как жить в сломавшемся мире > (вот,  например,  одно из писем ( https://mailchi.mp/getkit.news/readnews ) ).  > «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) *Денис Дмитриев* * Телеграм * Фб * Твиттер * Вк * Ок * В закладки * Напишите нам",
				},
			},
			{
				Title:    "Защитников «Азовстали» доставили в бывшую колонию под Донецком.  Всего с завода вывезли 959 украинских военных",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "08:05, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Украинских военных,  вышедших с территории завода «Азовсталь» в Мариуполе,  доставили в бывшую исправительную колонию в селе Еленовка под Донецком,  находящуюся под контролем военных России и самопровозглашенной ДНР.  Об этом сообщает агентство Reuters со ссылкой на очевидцев.  Завод «Азовсталь» был последним местом в оккупированном Мариуполе,  где укрывались украинские военнослужащие.  Вместе с ними было,  как сообщалось,  несколько сотен гражданских.  Как минимум часть из них эвакуировали в начале мая.  Вечером 16 мая с территории «Азовстали» вышли более 260 военных.  На следующий день с завода выехали ( https://meduza.io/news/2022/05/17/smi-soobschili-o-vyezde-s-azovstali-novoy-gruppy-ukrainskih-voennyh ) еще семь автобусов с военными.  В Минобороны РФ заявляют,  что украинские военные сдались в плен.  По данным ( https://t.me/mod_russia/15837 ) ведомства,  всего с завода вывезены 959 военных,  из них 51 раненый доставлен в больницу Новоазовска в ДНР.  Украинские власти пообещали вернуть защитников «Азовстали» в рамках обмена.  В Следственном комитете РФ заявили ( https://tass.ru/proisshestviya/14651461 ) ,  что «допросят сдавшихся боевиков,  укрывавшихся на комбинате „Азовсталь“ в Мариуполе».  В свою очередь спикер Госдумы Вячеслав Володин предложил ( https://meduza.io/news/2022/05/17/natsistskie-prestupniki-ne-dolzhny-podlezhat-obmenu-volodin-predlozhil-zapretit-obmen-zaschitnikov-azovstali-na-rossiyskih-plennyh ) запретить обмен бойцов «Азова»,  заявив,  что они «нацистские преступники» и что их нужно судить.  Читайте также Украинских военных начали вывозить с «Азовстали».  Генштаб ВСУ говорит,  что их передадут Киеву «через процедуру обмена» Россия это пока не подтвердила ----------------------------------------------------------------------------------------------------------------------------------------------------- 2 дня назад ( /feature/2022/05/16/sudya-po-vsemu-ukrainskih-voennyh-nachali-evakuirovat-s-azovstali-polk-azov-zayavil-chto-vypolnil-prikaz-i-nadeetsya-na-podderzhku-ukraintsev ) Читайте также Украинских военных начали вывозить с «Азовстали».  Генштаб ВСУ говорит,  что их передадут Киеву «через процедуру обмена» Россия это пока не подтвердила ----------------------------------------------------------------------------------------------------------------------------------------------------- 2 дня назад ( /feature/2022/05/16/sudya-po-vsemu-ukrainskih-voennyh-nachali-evakuirovat-s-azovstali-polk-azov-zayavil-chto-vypolnil-prikaz-i-nadeetsya-na-podderzhku-ukraintsev ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Швеция и Финляндия подали заявки на вступление в НАТО",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "06:16, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Генеральный секретарь НАТО Йенс Столтенберг принял 18 мая в Брюсселе у послов Финляндии и Швеции заявки на членство этих стран в альянсе.  Обе страны решили отказаться от нейтрального статуса и присоединиться к военному альянсу на фоне российского вторжения в Украину.  17 мая заявку Швеции на вступление в НАТО подписала министр иностранных дел страны Анн Линде.  В тот же день парламент Финляндии одобрил решение о присоединении к альянсу 188 голосами «за» и восьмью «против».  Заявки стран рассмотрит Совет НАТО,  после чего соглашение о членстве Швеции и Финляндии должны ратифицировать все 30 стран-участниц альянса.  Власти Турции уже выдвинули список условий,  на которых они согласятся одобрить вступление Швеции и Финляндии в НАТО.  По данным ( https://meduza.io/news/2022/05/17/turtsiya-vydvinula-usloviya-dlya-vstupleniya-shvetsii-i-finlyandii-v-nato-v-spiske-postavka-novyh-istrebiteley-i-priznanie-rabochey-partii-kurdistana-terroristami ) Bloomberg, Турция требует,  чтобы Швеция и Финляндия признали Рабочую партию Курдистана и связанные с ней структуры террористическими организациями.  Также страна рассчитывает получить новую партию истребителей F-16 и вновь попасть в программу поставок истребителей F-35.  Читайте также ------------- * Финляндия и Швеция собираются вступить в НАТО из-за войны в Украине.  Путин и тут просчитался? И что теперь ждет Россию? Объясняет политолог Кирилл Шамиев ( /feature/2022/04/21/finlyandiya-i-shvetsiya-sobirayutsya-vstupit-v-nato-iz-za-voyny-v-ukraine-putin-i-tut-proschitalsya-i-chto-teper-zhdet-rossiyu ) * Швеция была нейтральной страной больше 200 лет,  Финляндия — больше 75.  Но из-за войны в Украине обе вступают в НАТО.  Что это значит для России? ( /episodes/2022/05/17/shvetsiya-byla-neytralnoy-stranoy-bolshe-200-let-finlyandiya-bolshe-75-no-iz-za-voyny-v-ukraine-obe-vstupayut-v-nato-chto-eto-znachit-dlya-rossii ) Главные события дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- час назад ( /live/2022/05/18/voyna ) Главные события дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- час назад ( /live/2022/05/18/voyna ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "В России вновь попытались поджечь военкомат.  «Коктейлями Молотова» забросали здание Щелковского военного комиссариата в Подмосковье",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "07:57, 18 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) В городе Щелково Московской области в ночь на 18 мая в здание военного комиссариата бросили два «коктейля Молотова»,  сообщают Baza ( https://t.me/bazabazon/11661 ) и «Острожно новости» ( https://t.me/ostorozhno_novosti/8377 ).  В результате поджога пострадали два кабинета здания,  в том числе — архив комиссариата.  На месте происшествия работают пожарные,  полиция и сотрудники Следственного комитета РФ.  Сообщается,  что идет розыск поджигателя,  но не уточняется,  есть ли у полиции конкретные подозреваемые.  По подсчетам Baza, с начала вторжения России в Украину это уже 12 попытка поджога военкомата в разных регионах РФ.  Все события этого дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- 3 часа назад ( /live/2022/05/18/voyna ) Все события этого дня Война Восемьдесят четвертый день.  Онлайн «Медузы» ------------------------------------------------- 3 часа назад ( /live/2022/05/18/voyna ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
			{
				Title:    "Bloomberg: США планируют после 25 мая заблокировать выплаты России по госдолгу",
				Subtitle: "",
				Article: entities.Article{
					Author: "",
					Date:   "19:22, 17 мая 2022",
					Text:   "Данное сообщение (материал) создано и (или) распространено иностранным средством массовой информации,  выполняющим функции иностранного агента,  и (или) российским юридическим лицом,  выполняющим функции иностранного агента.  Нам нужна ваша помощь.  Пожалуйста,  поддержите «Медузу».  ( https://mdza.io/iRpGL_50KH4 ) Соединенные Штаты собираются полностью лишить Россию возможности платить по государственным облигациям,  сообщает Bloomberg со ссылкой на источники.  По данным собеседников агентства,  Минфин США не собирается продлевать лицензию,  позволявшую России проводить выплаты,  срок действия которой истекает 25 мая.  Некоторые чиновники Минфина США,  пишет Bloomberg, выступали за то,  чтобы продлить лицензию и дать,  таким образом,  России возможность тратить золотовалютные запасы на выплаты по госдолгу,  а не на войну.  Однако в администрации президента США Джо Байдена склонились к более жесткому сценарию.  Окончательное решение о сроках действия лицензии пока не принято.  27 мая России предстоит расплатиться по двум выпускам еврооблигаций,  отмечает агентство,  в случае,  если лицензия не будет продлена,  то РФ окажется на грани дефолта.  В начале мая Россия сумела избежать ( https://meduza.io/news/2022/05/03/bloomberg-rossiya-izbezhala-defolta-po-evrobondam ) дефолта по внешнему долгу в иностранной валюте,  проведя выплаты в долларах по двум выпускам суверенных еврооблигаций,  срок погашения которых закончился 4 апреля.  До этого РФ попыталась ( https://meduza.io/news/2022/04/06/minfin-rf-ne-smog-rasplatitsya-po-evrobondam-v-dollarah-i-vpervye-perechislil-vyplaty-v-rublyah-teper-rossii-mozhet-grozit-defolt ) провести выплаты по займам в рублях,  но агентство Moodyʼs заявило,  что это является изменением условий выпуска облигаций и может считаться дефолтом.  Что происходит с российской экономикой -------------------------------------- * Сколько еще проживет российская экономика? А потом будет как в девяностых? Что тогда будет делать правительство? Отвечает экономист Олег Буклемишев.  Во время дефолта 1998 года он работал в Минфине ( /feature/2022/03/10/skolko-esche-prozhivet-rossiyskaya-ekonomika-a-potom-budet-kak-v-devyanostyh-chto-togda-budet-delat-pravitelstvo ) * Курс рубля растет уже полтора месяца (хотя экономика под санкциями).  Что происходит? И когда начнется новое падение? ( /feature/2022/05/13/kurs-rublya-rastet-uzhe-poltora-mesyatsa-hotya-ekonomika-pod-sanktsiyami-chto-proishodit-i-kogda-nachnetsya-novoe-padenie ) * Как изменилась российская экономика за полтора месяца войны? И чего ждать дальше? «Это экономика выживания.  В лучшем случае экономика стагнации» ( /feature/2022/04/13/kak-izmenilas-rossiyskaya-ekonomika-za-poltora-mesyatsa-voyny-i-chego-zhdat-dalshe ) «Медуза».  Работаем 24/7.  И только в интересах читателей Нам срочно нужна ваша поддержка --------------------------------------------------------------------------------------- Хочу помочь ( https://mdza.io/WP4YuV4xTEM ) * Телеграм * Фб * Твиттер * Вк * Ок * Напишите нам",
				},
			},
		}

		p := tea.NewProgram(cli.PopulateGeneralNewsModel(news))
		if err := p.Start(); err != nil {
			logger.Printf("error implementing cli mode: %v", err)
			os.Exit(1)
		}

		return
	} else if mode == "http" {

		return
	}

	logger.Fatalf("cannot implement %s mode", mode)
}
