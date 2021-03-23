import logo from './logo.svg';
import './App.css';
import UploadForm from './component/uploadform/UploadForm'
function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <UploadForm></UploadForm>
      </header>
    </div>
  );
}

export default App;
