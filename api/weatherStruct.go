package api

type WeatherForecast struct {
	Lat            float64 `json:"lat"`
	Lon            float64 `json:"lon"`
	Timezone       string  `json:"timezone"`
	TimezoneOffset int     `json:"timezone_offset"`
	Daily          []struct {
		Dt        int     `json:"dt"`
		Sunrise   int     `json:"sunrise"`
		Sunset    int     `json:"sunset"`
		Moonrise  int     `json:"moonrise"`
		Moonset   int     `json:"moonset"`
		MoonPhase float64 `json:"moon_phase"`
		Temp      struct {
			Day   float64 `json:"day"`
			Min   float64 `json:"min"`
			Max   float64 `json:"max"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"temp"`
		FeelsLike struct {
			Day   float64 `json:"day"`
			Night float64 `json:"night"`
			Eve   float64 `json:"eve"`
			Morn  float64 `json:"morn"`
		} `json:"feels_like"`
		Pressure  int     `json:"pressure"`
		Humidity  int     `json:"humidity"`
		DewPoint  float64 `json:"dew_point"`
		WindSpeed float64 `json:"wind_speed"`
		WindDeg   int     `json:"wind_deg"`
		WindGust  float64 `json:"wind_gust"`
		Weather   []struct {
			Id          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Clouds int     `json:"clouds"`
		Pop    float64 `json:"pop"`
		Snow   float64 `json:"snow,omitempty"`
		Uvi    float64 `json:"uvi"`
		Rain   float64 `json:"rain,omitempty"`
	} `json:"daily"`
}

type CoordinatesCity []City

type City struct {
	Name       string `json:"name"`
	LocalNames struct {
		Fa          string `json:"fa"`
		Ab          string `json:"ab"`
		Ur          string `json:"ur"`
		Jv          string `json:"jv"`
		Na          string `json:"na"`
		Sw          string `json:"sw"`
		Sq          string `json:"sq"`
		Bn          string `json:"bn"`
		Ce          string `json:"ce"`
		Ga          string `json:"ga"`
		No          string `json:"no"`
		Hy          string `json:"hy"`
		Mg          string `json:"mg"`
		Sl          string `json:"sl"`
		Af          string `json:"af"`
		Ay          string `json:"ay"`
		Bs          string `json:"bs"`
		Kg          string `json:"kg"`
		Mn          string `json:"mn"`
		Fy          string `json:"fy"`
		Ml          string `json:"ml"`
		Lt          string `json:"lt"`
		He          string `json:"he"`
		So          string `json:"so"`
		Eo          string `json:"eo"`
		Sm          string `json:"sm"`
		Dz          string `json:"dz"`
		Mk          string `json:"mk"`
		Ch          string `json:"ch"`
		Es          string `json:"es"`
		FeatureName string `json:"feature_name"`
		St          string `json:"st"`
		My          string `json:"my"`
		Hi          string `json:"hi"`
		Gl          string `json:"gl"`
		Lv          string `json:"lv"`
		Ln          string `json:"ln"`
		Oc          string `json:"oc"`
		Tr          string `json:"tr"`
		Wa          string `json:"wa"`
		Sc          string `json:"sc"`
		Ro          string `json:"ro"`
		Mt          string `json:"mt"`
		Ps          string `json:"ps"`
		Io          string `json:"io"`
		De          string `json:"de"`
		Cy          string `json:"cy"`
		Kn          string `json:"kn"`
		Is          string `json:"is"`
		Ascii       string `json:"ascii"`
		Cs          string `json:"cs"`
		Su          string `json:"su"`
		Ht          string `json:"ht"`
		Wo          string `json:"wo"`
		Ar          string `json:"ar"`
		Te          string `json:"te"`
		Et          string `json:"et"`
		Yo          string `json:"yo"`
		Ss          string `json:"ss"`
		Nb          string `json:"nb"`
		Bi          string `json:"bi"`
		Co          string `json:"co"`
		Th          string `json:"th"`
		An          string `json:"an"`
		Uk          string `json:"uk"`
		Zu          string `json:"zu"`
		Fo          string `json:"fo"`
		Sg          string `json:"sg"`
		Pt          string `json:"pt"`
		Sh          string `json:"sh"`
		Uz          string `json:"uz"`
		Yi          string `json:"yi"`
		Fi          string `json:"fi"`
		El          string `json:"el"`
		Lg          string `json:"lg"`
		Ja          string `json:"ja"`
		Li          string `json:"li"`
		En          string `json:"en"`
		Av          string `json:"av"`
		Ie          string `json:"ie"`
		Dv          string `json:"dv"`
		Vo          string `json:"vo"`
		Gv          string `json:"gv"`
		Cv          string `json:"cv"`
		Eu          string `json:"eu"`
		Be          string `json:"be"`
		Gd          string `json:"gd"`
		Sr          string `json:"sr"`
		Ku          string `json:"ku"`
		Mi          string `json:"mi"`
		Tt          string `json:"tt"`
		Fr          string `json:"fr"`
		Ta          string `json:"ta"`
		Tg          string `json:"tg"`
		Kl          string `json:"kl"`
		Kv          string `json:"kv"`
		Qu          string `json:"qu"`
		Nn          string `json:"nn"`
		Ru          string `json:"ru"`
		Hu          string `json:"hu"`
		Ug          string `json:"ug"`
		Vi          string `json:"vi"`
		Gn          string `json:"gn"`
		Bo          string `json:"bo"`
		Os          string `json:"os"`
		Tl          string `json:"tl"`
		Az          string `json:"az"`
		Pl          string `json:"pl"`
		Cu          string `json:"cu"`
		Ka          string `json:"ka"`
		Ky          string `json:"ky"`
		Kw          string `json:"kw"`
		Nl          string `json:"nl"`
		Ba          string `json:"ba"`
		Id          string `json:"id"`
		Se          string `json:"se"`
		Za          string `json:"za"`
		Ca          string `json:"ca"`
		Zh          string `json:"zh"`
		Ak          string `json:"ak"`
		Da          string `json:"da"`
		Ms          string `json:"ms"`
		La          string `json:"la"`
		Sk          string `json:"sk"`
		Hr          string `json:"hr"`
		Tk          string `json:"tk"`
		Iu          string `json:"iu"`
		Ty          string `json:"ty"`
		Ia          string `json:"ia"`
		Br          string `json:"br"`
		Bg          string `json:"bg"`
		Sv          string `json:"sv"`
		Am          string `json:"am"`
		Ko          string `json:"ko"`
		Mr          string `json:"mr"`
		Kk          string `json:"kk"`
		It          string `json:"it"`
	} `json:"local_names"`
	Lat     float64 `json:"lat"`
	Lon     float64 `json:"lon"`
	Country string  `json:"country"`
	State   string  `json:"state"`
}
