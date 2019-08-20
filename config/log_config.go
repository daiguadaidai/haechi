package config

func LogDefautConfig() string {
	return `
        <seelog type="sync">
        	<outputs formatid="main">
                <console />
        	</outputs>
            <formats>
                <format id="main" format="%Date %Time [%Level] %File:%Line %Msg%n"/>
            </formats>
        </seelog>
    `
}
