import config from './util/config'

// const log = console.log

function Com(){
  
  const postData = () =>{
    fetch(config.APP_URL,{
      method: "POST",
      headers: {
        "Content-Type":"application/json"
      },
      body: JSON.stringify({
        "user":"yale918",
        "install":"redis"
      })
    })
    .then(res=>res.json())
    .then(data=>{
      console.log(data)
    })
    .catch(err=>{
      console.log(err)
    })
  }

  return (
    <>
      <input 
        type="button"
        value="redis"
        onClick={()=>{
          postData()
        }} 
      />
    </>
  )
}

export default Com