<template>
  <div class="container">
        <div class="row">
            <div class="col">
                <h1 class="mt-5">Test microservices</h1>
                <hr>
                <a id="brokerBtn" class="btn btn-outline-secondary" href="javascript:void(0);">Test Broker</a>
                <a id="authBrokerBtn" class="btn btn-outline-secondary" href="javascript:void(0);">Test Auth</a>
                <a id="logBtn" class="btn btn-outline-secondary" href="javascript:void(0);">Test Log</a>
                <a id="mailBtn" class="btn btn-outline-secondary" href="javascript:void(0);">Test Mail</a>

                <div id="output" class="mt-5" style="outline: 1px solid silver; padding: 2em;">
                    <span class="text-muted">Output shows here...</span>
                </div>
            </div>
        </div>
        <div class="row">
            <div class="col">
                <h4 class="mt-5">Sent</h4>
                <div class="mt-1" style="outline: 1px solid silver; padding: 2em;">
                    <pre id="payload"><span class="text-muted">Nothing sent yet...</span></pre>
                </div>
            </div>
            <div class="col">
                <h4 class="mt-5">Received</h4>
                <div class="mt-1" style="outline: 1px solid silver; padding: 2em;">
                    <pre id="received"><span class="text-muted">Nothing received yet...</span></pre>
                </div>
            </div>
        </div>
    </div> 
</template>

<script>
export default {
  name: 'MyContents',
}
</script>

<style>
</style>
