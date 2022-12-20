<template>
  <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-5">Tech Trend 📈</h1>
                <hr>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <div id="brokerBtn" class="btn btn-outline-secondary" @click="testBroker" >Test Broker</div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <div id="languageBtn" class="btn btn-outline-secondary" @click="getLanguages" >Test Language</div>
                <div id="webframeworkBtn" class="btn btn-outline-secondary" @click="getWebframeworks" >Test Webframework</div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <h2 class="mt-5">Programming languages</h2>
                <div class="mt-1" style="padding: 2em;">
                    // 차트
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <h2 class="mt-5">Programming languages</h2>
                <div class="mt-1" style="padding: 2em;">
                    // 차트
                </div>
            </div>
        </div>
    </div> 
</template>

<script>
import axios from 'axios';

export default {
  name: 'MyContents',
  methods: {
    testBroker() {
        axios.post('http://localhost:8080', {})
            .then(res => {
                console.log(res);
            });
    },
    getLanguages() {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        
        const body = {
            action: "survey/languages"
        };

        axios.post('http://localhost:8080/handle', body, {
            headers: headers
        })
            .then(res => {
                console.log(res);
            });
    },
    getWebframeworks() {
        const headers = new Headers();
        headers.append("Content-Type", "application/json");
        
        const body = {
            action: "survey/webframeworks"
        };

        axios.post('http://localhost:8080/handle', body, {
            headers: headers
        })
            .then(res => {
                console.log(res);
            });
    }
  }
}
</script>

<style>
h1, h2{
    color: #2c3e50;
    font-weight: 400;
}
h2 {
    font-size: 1.2em;
}
</style>
