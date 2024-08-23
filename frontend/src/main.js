import './style.css';
import './app.css';
import {Greet, GetTodayUsage} from '../wailsjs/go/main/App';

document.querySelector('#app').innerHTML = `
    <div class="table-container">
        <div class="header">
            <h1>Siti Visitati</h1>
        </div>
        <table id="table">
            <tr>
                <th>Nome Sito</th>
                <th>Data</th>
                <th>Tempo Passato</th>
            </tr>
        </table>
    </div>
`;

window.getusage = function () {
	try {
		GetTodayUsage()
			.then((res) => {
				const table = document.getElementById("table");
				for (let program of res) {
					let prg = document.createElement("td")
					let day = document.createElement("td")
					let time = document.createElement("td")

					prg.innerText = program.Program;
					day.innerText = program.Day;
					time.innerText = program.Time;

					let tr = document.createElement("tr")

					tr.appendChild(prg);
					tr.appendChild(day);
					tr.appendChild(time);

					table.appendChild(tr);
				}
			})
			.catch((err) => {
				console.error(err);
			});
	} catch (err) {
		console.error(err);
	}
}

getusage();
