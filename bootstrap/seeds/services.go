package seeds

import (
	"log"

	"github.com/gabrielfmcoelho/platform-core/domain"
	"gorm.io/gorm"
)

// SeedServices cria os serviços padrão (Resistracker, ...) se não existirem
func SeedServices(db *gorm.DB) error {
	var count int64
	if err := db.Model(&domain.Service{}).Count(&count).Error; err != nil {
		return err
	}

	if count == 0 {
		services := []domain.Service{
			{
				MarketingName: "Resistracker",
				Name:          "Resistracker",
				Description:   "Acompanhamento de registros",
				AppUrl:        "https://ipsdatalab.shinyapps.io/resistracker_sao_marco/",
				IconUrl:       "https://resistracker.solude.tech/favicon.ico",
				ScreenshotUrl: "https://static.wixstatic.com/media/d0e9f8_40e18c9c8238461fb62e08b04a30b1e4~mv2.png/v1/crop/x_286,y_0,w_946,h_779/fill/w_454,h_374,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/Design%20sem%20nome%20(2).png",
				TagLine:       "Gestão Inteligente de Resistência Bacteriana",
				Benefits:      "Redução de 40% no tempo de identificação de padrões de resistência;Aumento de 60% na eficácia do tratamento inicial;Economia de 30% nos custos com antibióticos",
				Features:      "Monitoramento em tempo real;Análise preditiva de resistência;Suporte à decisão clínica;Relatórios personalizados",
				Tags:          "IA;Microbiologia;Antibióticos",
				LastUpdate:    "2021-09-01",
				Status:        "Online",
				Version:       "1.0.0",
				Price:         0.00,
				IsMarketing:   true,
			},
			{
				MarketingName: "Exame Extractor",
				Name:          "Exame Extractor",
				Description:   "",
				AppUrl:        "",
				IconUrl:       "https://static.wixstatic.com/media/d0e9f8_0e5a688098944276ace8935455185ed2~mv2.png/v1/fill/w_31,h_31,al_c,lg_1,q_85,enc_avif,quality_auto/brain-circuit.png",
				ScreenshotUrl: "https://static.wixstatic.com/media/d0e9f8_08500bf86671472aa205d74feafd1e76~mv2.png/v1/fill/w_470,h_373,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/Design%20sem%20nome%20(1).png",
				TagLine:       "Automação Inteligente de Dados Laboratoriais",
				Benefits:      "",
				Features:      "Extração automática de dados;Integração com sistemas;Validação inteligente;Dashboards em tempo real",
				Tags:          "Laboratório;Automação;Integração",
				LastUpdate:    "",
				Status:        "",
				Version:       "",
				Price:         0.00,
				IsMarketing:   true,
			},
			{
				MarketingName: "Solude Bioanalytics",
				Name:          "Solude Bioanalytics",
				Description:   "",
				AppUrl:        "",
				IconUrl:       "https://static.wixstatic.com/media/d0e9f8_fcadc8b7e9504858962cdca5a88b0ac1~mv2.png/v1/fill/w_31,h_31,al_c,lg_1,q_85,enc_avif,quality_auto/stethoscope.png",
				ScreenshotUrl: "https://static.wixstatic.com/media/d0e9f8_e37c15d9556c4519a22dd3c6d5298b8b~mv2.png/v1/crop/x_0,y_0,w_980,h_779/fill/w_470,h_374,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/Design%20sem%20nome.png",
				TagLine:       "Análise Avançada para Decisões Clínicas",
				Benefits:      "",
				Features:      "Análise preditiva de resultados;Alertas inteligentes;Visualização interativa de dados;Relatórios personalizados",
				Tags:          "IA;Analytics;Decisão Clínica",
				LastUpdate:    "",
				Status:        "",
				Version:       "",
				Price:         0.00,
				IsMarketing:   true,
			},
			{
				MarketingName: "Pharmawatch",
				Name:          "Pharmawatch",
				Description:   "",
				AppUrl:        "",
				IconUrl:       "https://static.wixstatic.com/media/d0e9f8_dad1207f17b44acaae97ac668c3e6d11~mv2.png/v1/fill/w_31,h_31,al_c,lg_1,q_85,enc_avif,quality_auto/chart-line.png",
				ScreenshotUrl: "https://static.wixstatic.com/media/d0e9f8_acae4852c6df489d931dc2921dc634ef~mv2.png/v1/fill/w_470,h_373,al_c,q_85,usm_0.66_1.00_0.01,enc_avif,quality_auto/Design%20sem%20nome%20(4).png",
				TagLine:       "Gestão Estratégica de Custos Farmacêuticos",
				Benefits:      "",
				Features:      "Análise de custos em tempo real;Gestão de estoque inteligente;Previsão de demanda;Relatórios financeiros automatizados",
				Tags:          "Farmácia;Custos;Gestão",
				LastUpdate:    "",
				Status:        "",
				Version:       "",
				Price:         0.00,
				IsMarketing:   true,
			},
		}

		if err := db.Create(&services).Error; err != nil {
			return err
		}

		log.Printf("[SeedServices] Criados %d serviços (Resistracker, ...)\n", len(services))
	}
	return nil
}
