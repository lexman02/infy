export default function Signup(){
    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Fill out your INFYnaut Application!</h1>
            </div>
            <br/>
            <div className="flex justify-center align-middle">
                <form>
                    <label for="username" className="text-neutral-50">Username:</label><br/>
                    <input type="text" id="username" name="username" className="hover:bg-neutral-400"></input><br/><br/>
                    <label for="password" className="text-neutral-50">Password:</label><br/>
                    <input type="password" id="password" name="password" className="hover:bg-neutral-400"></input><br/><br/>
                    <label for="email" className="text-neutral-50">Email:</label><br/>
                    <input type="text" id="email" name="email" className="hover:bg-neutral-400"></input><br/><br/>
                    <input type="submit" value="Submit" className="bg-violet-900 text-neutral-50 px-4 py-2 rounded-lg hover:bg-violet-950" />
                </form>
            </div>
        </div>
    )
}
/*
const form = document.querySelector('form')
form.addEventListener('submit', (e) => {
  e.preventDefault()
  const formData = new FormData(e.target)
  const json = JSON.stringify(Object.fromEntries(formData));
})
*/