const ctxComplexity = document.getElementById('complexityChart').getContext('2d');
const bar = document.getElementById('barChart').getContext('2d');

const chart = new Chart(bar, {
    type: 'bar',
    data: {
        labels: ['Iteratif', 'Rekursif'],
        datasets: [{
            label: 'Waktu Eksekusi Terakhir (ms)',
            data: [0, 0],
            backgroundColor: ['#007bff', '#28a745'],
            borderRadius: 5,
        }]
    },
    options: {
        responsive: true,
        plugins: { legend: { display: false } },
        scales: { y: { beginAtZero: true } }
    }
});


const lineChart = new Chart(ctxComplexity, {
  type: 'line',
  data: {
    datasets: [
      {
        label: 'Iteratif',
        data: [],
        borderColor: '#007bff', 
        tension: 0.3
      },
      {
        label: 'Rekursif',
        data: [],
        borderColor: '#28a745',
        tension: 0.3
      }
    ]
  },
  options: {
    scales: {
      x: {
        type: 'linear',
        beginAtZero: true,
        title: {
          display: true,
          text: 'Banyak Bilangan (n)'
        }
      },
      y: {
        type: 'linear',
        beginAtZero: true,
        title: {
          display: true,
          text: 'Waktu Eksekusi (ms)'
        }
      }
    }
  }
});

function processMedian(metode) {
    if (metode === 'iteratif') {
        runIterative();
    } else {
        runRecursive();
    }
}

let iteratifTime = 0;
let rekursifTime = 0;

async function runIterative() {

    const input = document.getElementById('numberInput').value;
    const outputElement = document.getElementById('iteratifOutput');
    const timeElement = document.getElementById('iteratifTime');

    const numbers = input.split(/[\s,]+/).map(Number).filter(n => !isNaN(n));
    console.log(numbers);
    if (numbers.length === 0) {
        outputElement.textContent = "Input tidak valid";
        return;
    }

    try {
        const response = await fetch('http://localhost:8080/proses', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                metode: 'iteratif',
                data: numbers
            })
        });

        const data = await response.json();

        outputElement.textContent = data.median;
        timeElement.textContent = data.executionTime.toFixed(2) + " ms";
        iteratifTime = parseFloat(data.executionTime);

        updateChart()
        updateLineChart(numbers.length, iteratifTime, 'iteratif');
    } catch {
        outputElement.textContent = "Error Server";
    }
}

async function runRecursive() {
    const input = document.getElementById('numberInput').value;
    const outputElement = document.getElementById('rekursifOutput');
    const timeElement = document.getElementById('rekursifTime');

    const numbers = input.split(/[\s,]+/).map(Number).filter(n => !isNaN(n));
    if (numbers.length === 0) {
        outputElement.textContent = "Input tidak valid";
        return;
    }


    try {
        const response = await fetch('http://localhost:8080/proses', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({
                metode: 'rekursif',
                data: numbers
            })
        });

        const data = await response.json();
  
        outputElement.textContent = data.median;
        timeElement.textContent = data.executionTime.toFixed(2) + " ms";
        rekursifTime = parseFloat(data.executionTime);

        updateChart()
        updateLineChart(numbers.length, rekursifTime, 'rekursif');
    } catch {
        outputElement.textContent = "Error Server";
    }
}

function updateChart() {
    chart.data.datasets[0].data = [iteratifTime, rekursifTime];
    chart.update();
}

function updateLineChart(n, time, method) {
    const datasetIndex = (method === 'iteratif') ? 0 : 1;

    lineChart.data.datasets[datasetIndex].data.push({ x: n, y: time });
    lineChart.data.datasets[datasetIndex].data.sort((a, b) => a.x - b.x);
    lineChart.update('none');
}


