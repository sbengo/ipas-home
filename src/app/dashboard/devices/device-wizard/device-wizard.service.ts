import { HttpClient } from '@angular/common/http';
import { Injectable, } from '@angular/core';
import 'rxjs/add/operator/map';
import { Observable } from 'rxjs/Observable';

declare var _:any;

@Injectable()
export class DeviceWizardService {

    constructor(public httpAPI: HttpClient) {
        console.log('Task Service created.', httpAPI);
    }
    
    createNewProduct(url: string, dev : any) {
     console.log(url, dev)
     return this.httpAPI.post(url,dev)
        .map((responseData) =>
            { console.log(responseData); return responseData}
        )
    };

    newDevice(platform: any, device: any) {
        console.log("DEVICEE", device);
        let dev =  JSON.parse(JSON.stringify(device))
        let pTags = platform.tags.map((element) => { return element.key+'='+element.value} )
        console.log(pTags)
        //Map platform tags as device tags only on snmpcollector engine
        
        let ie = dev.engine.findIndex((engine) => engine.name === 'snmpcollector')  
        let iparam = dev.engine[ie].params.findIndex((param) => param.key === 'DEVICE_EXTRATAG_VALUES')
    
        dev.engine[ie].params[iparam].value = dev.engine[ie].params[iparam].value.split(',').concat(pTags)
    
        //Prepare structure to be sent:
        let finalForm = {'platform' : platform, 'devices': [dev]}

        console.log(finalForm);

      }


    getPlatformEngines(url:string) {
       return this.httpAPI.get(url)
    }

    uploadFiles(url,formGroup: any, files : any) {
        const formData: any = new FormData();        
        formData.append('Msg','MyCustomMessage' );
        //All engines
        for (let i in files.gather) {
            //All config of iengine
            for (let j in  files.gather[i].config){
                //All available configs...
                for (let p in files.gather[i].config[j]) {
                    //Generate output dir: /products/<PRODUCT_NAME>/<DIR>/<SOURCE>
                    let rootDir =  '/products/'
                    let product =  formGroup.product+'/'
                    let dir = formGroup['gather'][i].config[j].dir+'/'
                    let source = formGroup['gather'][i].config[j].config[p].source
                    formData.append("CommitFile", files.gather[i].config[j][p][0], rootDir+product+dir+source);
                }
            }
        }
        return this.httpAPI.post(url, formData)
    }

    retrieveYAML(url: string) {
        console.log(url)
        return this.httpAPI.get(url, {'responseType': 'text'})
           .map((responseData) =>
           { console.log(responseData); return responseData}
           )
       };
}