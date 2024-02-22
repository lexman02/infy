export default function Login(){
    return (
        <div className="bg-neutral-900 rounded-lg w-80 h-80 mx-auto text-center">
            <div className="text-neutral-50 flex justify-center content-center align-middle">
                <h1>Welcome back INFYnaut!</h1>
            </div>
            <br/>
            <div className="flex justify-center content-center align-middle">
                <form>
                    <label for="username" className="text-neutral-50">Username:</label><br/>
                    <input type="text" id="username" name="username"></input><br/><br/>
                    <label for="password" className="text-neutral-50">Password:</label><br/>
                    <input type="text" id="password" name="password"></input><br/><br/>
                    <input type="submit" value="Log In" className="bg-violet-900 text-neutral-50" />
                </form>
            </div>
        </div>
    )
}