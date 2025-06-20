package giga

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSummary(t *testing.T) {
	cli, err := NewGigaClient(context.Background(), "MTZmODdlMmUtOTJiNy00NzNlLWE0ODItMWY2OWU4ZjhmNjNlOmUwYWI5NDgwLTFjZjUtNDYzMi1hZWIxLTM0OWE0MmE1NDJiNg==")
	assert.NoError(t, err)

	txt := `Евгений Акпаев: @vladimir_nadulich раз тут пятница, скажи, у вас там производство соли в Мозыре сломалось что-ли? Год назад 100 кг брал за 1600, а сейчас 4000 минимум
	Vladimir Nadulich: Брал недавно кг за 0.78 коп кажется. Переводи
		Евгений Акпаев: Не, я о таблетированной
	Кирилл Черненко: Пажжи, на кой хрен тебе 100кг соли? Это среднестатистическому человеку хватит примерно на две жизни
		John Doe: Совместные закупки
		Евгений Акпаев: Для системы водоочистки
			Кирилл Черненко: Я определено чего то не знаю о системе водоочистки
				Евгений Акпаев: Она используются для рассола, который используется для умягчения воды и регенерации ионообменной смолы (которая нужна для удаления примесей железа) в системах водоподготовки. Часто используется в частных домах для очистки воды из скважин и в малых дозах - в посудомойках для умягчения (шоб белого налета не было)
					Кирилл Черненко: Ясненько.Ну у нас тут в целом вода очень мягкая, так что задачи умягчения в городе как концепции не стоит.А регенерация это прикольно
						Евгений Акпаев: Умягчение - это сайдэффект. Основное назначение - это очистка от железа при аэрации
							Vladimir Nadulich: зачем исправлять всю воду в доме?
								Евгений Акпаев: Не ушатыать электро и сантехнику
									Vladimir Nadulich: хке... чем?пусть это просто бойлер... нагревательному элементу есть до этого дело?
										Евгений Акпаев: Посудомойка, стиралка, вся сантехника, котел, бойлер, все трубы в доме, гидро аккумулятор, фильтра на питьевую воду. Это ещё далеко не все, что ушатывается
			Г🐈рри: а откуда ты воду добываешь? У тебя там скважина, или какая-то централизованная подача?
				Кирилл Черненко: При централизованной водоподготовка уже проведена, там как максимум фильтр тонкой очистки впаять остаётся.)
					Евгений Акпаев: Далекооооо не такая
						Кирилл Черненко: Я имею ввиду что для централизованной системы подачи воды не нужна своя станция водоподготовки в любом случае
							Евгений Акпаев: Хорошо бы, что-бы была. Но  такая дорогая не нужна (мне в 140 обошлась)
								Г🐈рри: Хочешь сказать, что централизованная подача воды у вас - это типа привелегия богатых?
									Евгений Акпаев: Ее просто нет. Где-то делают общую на коттеджный поселок, но ну его нафиг
								Кирилл Черненко: Да там достаточно обычного вводного узла, типа фильтр магистральной очистки, фильтр тонкой очистки, ну и если питьевую хочется ещё колбы впаять
				Евгений Акпаев: Скважина. Даже если бы была центральная, то все равно скважину прорыл бы. Не зависишь ни от кого и вода чище. Но денег стоит
					Г🐈рри: а глубина какая?
						Евгений Акпаев: У нас тут 70
							Г🐈рри: А если к ней ухом прильнуть, - голоса слышно оттудова?
								Кирилл Черненко: Рабочих которые скважину копали?
								Евгений Акпаев: Не😁
								Никита Авдеев: Мне кажется ад находится чуть дальше 70-ти метров
					Vladimir Nadulich: На даче соседи до австралии пробурили - и все равно болото (пахнет сероводородом)Другому соседу "прокачается" сказали работяги... прошло более 30 лет - не прокачалась ))Так что не всегда своя скважина лучше.За то (как мне показалось) волосы на моей лысой голове более шелковистые от сероводорода после душа))
						Г🐈рри: ванны тоже принимаешь?
							Vladimir Nadulich: не, ванны нету (это дача)
								Г🐈рри: Надо купить. И принимай.
						Евгений Акпаев: Артезианская (глубокая), часто, может быть гораздо более грязная, чем в верхних слоях (зато дебе больше). А сероводород вообще легко устраняется
							Vladimir Nadulich: хммм, как?
								Евгений Акпаев: Угольными фильтрами. BB10/BB20
									Евгений Акпаев: И ими не только сероводород, а вообще почти любые запахи
									Vladimir Nadulich: нужно попробоватьу меня водопотребление малое совсем в доме, бытовой техники нет
										Евгений Акпаев: Вот такой возьми (именно 0.6). Почувствуешь сильный качественный скачок
											Vladimir Nadulich: весной попробую) так то мы ее все равно не пьем
						Кирилл Черненко: А нехрен копать без геологии
							Vladimir Nadulich: шта? какие варианты? отступите от реки более километра?что такое геология? разведка бурением?
								Кирилл Черненко: Геологические изысканияЛюдей зовёшь, они у тебя копают в глубь и рассказывают где у тебя чё залегает
									Vladimir Nadulich: копать, чтобы потом пробудить 50мм трубу?Скважина стоит 10$ за метр. Вызвал, сделали, уехали. Копейки. На 100+мм в 2-3 раза дороже.Разведка - это та же скважина, просто еще одна попытка.
										Кирилл Черненко: Разведка все равно уже должна была быть до строительства дома, так что обычно к моменту скважины она так или иначе есть
											Vladimir Nadulich: купил дом в деревневот тебе река, вот конюшня, вот пьяный михалычОн вдвшник, вот он и служил в разведке.
												Никита Авдеев: В твоем доме в деревне?
													Vladimir Nadulich: Да, под дачу. А квартиры в городе.
													Кирилл Черненко: Это несущий Михалыч
													Vladimir Nadulich: а михалыч... это образ, признак, легенда, этнос ))
												Кирилл Черненко: И через два года тебе на голову упала твоя груда пиломатериалов сложенная в домоподобную конструкцию но по факту являющуюся сараем.Классика
													Vladimir Nadulich: Это когда ты строишься.У меня знакомые под молодечно в глухущей деревне купили дом - ему 100 лет. Венцы еще не сгнили. Открытые бревна снаружи. Кайфуют от "старины" и от глуши.Чему там уже проседать в почве?
														Кирилл Черненко: Если дому сто лет, то он сам уже как почва примерно.У них не такой срок эксплуатации.)
						Евгений Акпаев: Сероводород, кстати, не болотом пахнет (болото - это чуть застойная вода). При растворенном сероводороде стоит сильный запах го..на😄Скважина должна постоянно прокачиваться и чем больше - тем чище там будет
					John Doe: Электричество:- Добрый вечер...
						Andrey Konev: Разбаловала вас техническая цивилизация
							John Doe: Нас?
						Vladimir Nadulich: Костер во дворе: да пошло оно нахер, посидим в темноте, посмотрим на комету. Наливай.
							Евгений Акпаев: У нас когда отключают, мы прямо с удовольствием вываливаемся на улицу и жарим сёмгу в барбекюшнице с шампиньонами для гриля.PS: аж сейчас захотелось 😊
						Евгений Акпаев: Генератор на крайний случай
							John Doe: Бензин / газ:- Спокойной ночи
								Евгений Акпаев: Этого добра у меня всегда навалом. Техники куча всякой
									John Doe: На аккумуляторах не советуешь?
										Евгений Акпаев: Их менять надо периодически, оберегать от холода, т.к. деградируют. Муторно и дорого. Проще тарахтелку завести ненадолго
											John Doe: Это ты про генератор электричества?
												Евгений Акпаев: Да
											Vladimir Nadulich: если нужно дежурное освещение, да еще и ночью - то тарахтелка не очень ))
												Евгений Акпаев: Для этого есть аккумуляторные фонари. Но их очень редко используем
										Vladimir Nadulich: в телефоне фонарик норм)))а тебе для чего? сколько киловатт?
					Aleksandr Shumakov: Плывун поймаешь)Или вода желтая будет
						Евгений Акпаев: Не на такой глубине. А что-бы жёлтая не была (или ещё какая) надо чистить
		Nikolay Verhovcev: у него семья большая
			John Doe: В поле каждый колосок?
		Andrey Konev: Современный человек в среднем съедает 5 кг соли в год
			John Doe: А ты?
				Andrey Konev: не считал. Надеюсь, что меньше
					John Doe: Так и не сможешь посчитать, блюда же покупные там почти везде она в составе
			Кирилл Черненко: Я какой-то не современный`
	result, _ := cli.GetSummary(txt)
	fmt.Println(result)
}

//
//func Test_GetCommitMsg(t *testing.T) {
//	c := gomock.NewController(t)
//	defer c.Finish()
//
//
//	t.Run("error create", func(t *testing.T) {/**/
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return nil, errors.New("error")
//		})
//		defer p.Reset()
//
//		cli, err := NewGigaClient(context.Background(), "111", "222")
//		assert.Nil(t, cli)
//		assert.EqualError(t, err, "newGigaClient error: error")
//	})
//	t.Run("auth error", func(t *testing.T) {
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return new(gigachat.Client), nil
//		})
//		defer p.Reset()
//
//		client := mock_giga.NewMockIGigaClient(c)
//		client.EXPECT().AuthWithContext(gomock.Any()).Return(errors.New("error"))
//
//		cli, _ := NewGigaClient(context.Background(), "111", "222")
//		cli.client = client
//
//		_, _, _, err := cli.GetSpamPercent("")
//		assert.EqualError(t, err, "auth error: error")
//	})
//	t.Run("req error", func(t *testing.T) {
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return new(gigachat.Client), nil
//		})
//		defer p.Reset()
//
//		client := mock_giga.NewMockIGigaClient(c)
//		client.EXPECT().AuthWithContext(gomock.Any()).Return(nil)
//		client.EXPECT().ChatWithContext(gomock.Any(), gomock.Any()).Return(nil, errors.New("error"))
//
//		cli, _ := NewGigaClient(context.Background(), "111", "222")
//		cli.client = client
//
//		_, _, _, err := cli.GetSpamPercent("tyuyu")
//		assert.EqualError(t, err, "request error: error")
//	})
//	t.Run("response does not contain data", func(t *testing.T) {
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return new(gigachat.Client), nil
//		})
//		defer p.Reset()
//
//		client := mock_giga.NewMockIGigaClient(c)
//		client.EXPECT().AuthWithContext(gomock.Any()).Return(nil)
//		client.EXPECT().ChatWithContext(gomock.Any(), gomock.Any()).Return(&gigachat.ChatResponse{}, nil)
//
//		cli, _ := NewGigaClient(context.Background(), "111", "222")
//		cli.client = client
//
//		_, _, _, err := cli.GetSpamPercent("ghgh")
//		assert.EqualError(t, err, "response does not contain data")
//	})
//	t.Run("diff is not defined", func(t *testing.T) {
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return new(gigachat.Client), nil
//		})
//		defer p.Reset()
//
//		client := mock_giga.NewMockIGigaClient(c)
//		client.EXPECT().AuthWithContext(gomock.Any()).Return(nil)
//
//		cli, _ := NewGigaClient(context.Background(), "111", "222")
//		cli.client = client
//
//		_, _, _, err := cli.GetSpamPercent("")
//		assert.EqualError(t, err, "message is not defined")
//	})
//	t.Run("pass", func(t *testing.T) {
//		p := gomonkey.ApplyFunc(gigachat.NewInsecureClient, func(clientId string, clientSecret string) (*gigachat.Client, error) {
//			return new(gigachat.Client), nil
//		})
//		defer p.Reset()
//
//		client := mock_giga.NewMockIGigaClient(c)
//		client.EXPECT().AuthWithContext(gomock.Any()).Return(nil)
//		client.EXPECT().ChatWithContext(gomock.Any(), gomock.Any()).Return(&gigachat.ChatResponse{
//			Choices: []gigachat.Choice{{Message: gigachat.Message{Content: "89|в сообщении фигурирует фраза про криптовалюту и заработок"}}},
//		}, nil)
//
//		cli, _ := NewGigaClient(context.Background(), "111", "222")
//		cli.client = client
//
//		s, perc, r, err := cli.GetSpamPercent("hjhj")
//		assert.NoError(t, err)
//		assert.Equal(t, "в сообщении фигурирует фраза про криптовалюту и заработок", r)
//		assert.True(t, s)
//		assert.Equal(t, perc, 89)
//	})
//}
