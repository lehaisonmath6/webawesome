import React from 'react'
 
export const URL_UPLOAD = 'http://127.0.0.1:8080/fileupload'

class UploadForm extends React.Component {
    constructor(props){
        super(props);
        this.state ={
            selectedFile:null,
            filePathDownload:''
        }
        this.onFileChange = this.onFileChange.bind(this)
        this.onFileUpload = this.onFileUpload.bind(this)
    }
    onFileChange(event){
        this.setState({selectedFile:event.target.files[0]});
    }
    onFileUpload(){
        const formData = new FormData();
        formData.append("myFile",this.state.selectedFile,this.state.selectedFile.name)
        fetch(URL_UPLOAD,{
            body :formData,
            method:'post'
        }).then(res => res.json()).then(data=>{
            console.log("data",data)
            this.setState({
                filePathDownload : data['url']
            })
        })
    }
    render(){
        const  fileData = ()=>{
            if(this.state.selectedFile){
                return (
                    <div>
                    <h2>File Details:</h2>
                    <p>File Name: {this.state.selectedFile.name} </p>
                    <p>File type: {this.state.selectedFile.type} </p>
                    <p>Last modified: {this.state.selectedFile.lastModifiedDate.toDateString()}</p>
                    <p>Download at: {this.state.filePathDownload} </p>
                </div>)
            }else{
                return (
                    <div>
                        <br></br>
                        <h4>Choose before Pressing the Upload button</h4>
                    </div>
                )
            }
        }

        return(
            <div>
                <h1>Upload file</h1>
                <div>
                    <input type='file' onChange={this.onFileChange} ></input>
                    <button onClick={this.onFileUpload}>Upload</button>
                </div>
                {fileData()}
            </div>
        )
    }

}
export default UploadForm;