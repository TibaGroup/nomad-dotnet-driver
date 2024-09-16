using System;
using System.Threading;
using System.Threading.Tasks;
using Nancy;
using Nancy.Hosting.Self;
using System.IO;
using System.Runtime;
using System;
using System.Collections.Generic;
using Newtonsoft.Json;
using System;
using System.IO;

namespace TestNomadTask
{
    class Program
    {
        static void Main(string[] args)
        {
            var port = Environment.GetEnvironmentVariable("NOMAD_HOST_PORT_http");
            var uri = new Uri("http://localhost:" + port);
            var hostConfig = new HostConfiguration { UrlReservations = new UrlReservations { CreateAutomatically = true } };
            using (var host = new NancyHost(hostConfig, uri))
            {
                host.Start();
                Console.WriteLine("NancyFX service running on " + uri);

                // Keep the application running indefinitely
                Task.Run(() => KeepRunning()).Wait();
            }
        }

        private static void KeepRunning()
        {
            while (true)
            {
                Thread.Sleep(Timeout.Infinite);
            }
        }
    }
    public class DictionaryConverter
    {
        public static Dictionary<string, string> ConvertToStringDictionary(Dictionary<string, object> original)
        {
            var stringDictionary = new Dictionary<string, string>();

            foreach (var kvp in original)
            {
                stringDictionary[kvp.Key] = kvp.Value?.ToString();
            }

            return stringDictionary;
        }
    }
   public class ConfigProperties
   {
       public bool System_GC_Server { get; set; }
       public int System_Threading_ThreadPool_MaxThreads { get; set; }
       public int System_Threading_ThreadPool_MinThreads { get; set; }
   }

   public class Framework
   {
       public string Name { get; set; }
       public string Version { get; set; }
   }

   public class RuntimeOptions
   {
       public ConfigProperties ConfigProperties { get; set; }
       public Framework Framework { get; set; }
       public string Tfm { get; set; }
   }

   public class ConfigFile
   {
       public RuntimeOptions RuntimeOptions { get; set; }
   }
    public class MyModule : NancyModule
    {
        public MyModule()
        {
            Get("/__health", _ =>
            {
                var response = new
                {
                    status = "ok",
                    datetime = DateTime.UtcNow
                };
                return Response.AsJson(response);
            });

            Get("/configs", _ =>
            {
                string text = File.ReadAllText("./local/TestNomadTask.runtimeconfig.json");
                return Response.AsText(text);
            });
        }
    }
}