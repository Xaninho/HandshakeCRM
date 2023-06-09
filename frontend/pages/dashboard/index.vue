<template>
    <main>
      <div class="card">
          <Chart type="bar" :data="chartData" :options="chartOptions" />
      </div>
      <div class="card card-container">
          <Chart type="pie" :data="chartData" :options="chartOptions" class="chart-container" />
          <Chart type="line" :data="chartData" :options="chartOptions" class="chart-container" />
      </div>
    </main>
    
  </template>
  
  <script lang='ts'>
  
  
  export default defineComponent({
    name: "HomeDashboard",
    data() {
      return {
         chartData: {
            labels: ['Q1', 'Q2', 'Q3', 'Q4'],
            datasets: [
                {
                    label: 'Sales',
                    data: [540, 325, 702, 620],
                    backgroundColor: ['rgba(255, 159, 64, 0.2)', 'rgba(75, 192, 192, 0.2)', 'rgba(54, 162, 235, 0.2)', 'rgba(153, 102, 255, 0.2)'],
                    borderColor: ['rgb(255, 159, 64)', 'rgb(75, 192, 192)', 'rgb(54, 162, 235)', 'rgb(153, 102, 255)'],
                    borderWidth: 1
                }
            ]
          },
          chartOptions: {
              scales: {
                  y: {
                      beginAtZero: true
                  }
              }
          }
      }
     
    }
  });
  
  
  
  </script>
  
  <style scoped>
    .card-container {
    display: flex;
    flex-wrap: wrap;
    justify-content: space-between;
  }
  
  .chart-container {
    flex: 1;
    margin-right: 1rem;
    margin-bottom: 1rem;
  }
  
  .chart-container:last-child {
    margin-right: 0;
  }
  </style>