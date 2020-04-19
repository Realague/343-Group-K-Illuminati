using System;
using System.Collections.Generic;
using System.Diagnostics;
using System.IO;
using System.IO.Compression;
using System.Linq;
using System.Net;
using System.Text;
using System.Threading.Tasks;
using System.Windows;
using System.Windows.Controls;
using System.Windows.Data;
using System.Windows.Documents;
using System.Windows.Input;
using System.Windows.Media;
using System.Windows.Media.Imaging;
using System.Windows.Navigation;
using System.Windows.Shapes;

namespace IlluminatiLauncher
{
    /// <summary>
    /// Logique d'interaction pour MainWindow.xaml
    /// </summary>
    public partial class MainWindow : Window
    {
        private string APIEndPoint = "http://localhost:3000/api";

        public MainWindow()
        {
            InitializeComponent();
            CheckForUpdates();
        }

        private void PlayButton_Click(object sender, RoutedEventArgs e)
        {
            Launcher.PlayGame();
            Environment.Exit(0);
        }

        private void CheckForUpdates()
        {
            PlayButton.IsEnabled = false;
            string filename = ".version";
            string version = GetServerVersion();

            if (version == "")
            {
                PlayButton.IsEnabled = true;
            }
            if (!File.Exists(filename))
            {
                using (StreamWriter writer = new StreamWriter(filename, true))
                {
                    writer.WriteLine(version);
                }
                Update();
            }
            else
            {
                string localVersion = GetLocalVersion();
                if (localVersion != version)
                {
                    Update();
                }
            }
            PlayButton.IsEnabled = true;
        }

        private void Update()
        {
            try
            {
                File.Delete(".version");
                Directory.Delete("Illuminati_Data", true);
                File.Delete("UnityPlayer.dll");
                File.Delete("Illuminati.exe");
            } catch
            {

            }

            using (var client = new WebClient())
            {
                client.DownloadFile("http://localhost:8080/Illuminati/Illuminati.zip", "Illuminati.zip");

          //      ZipFile.ExtractToDirectory("Illuminati.zip", ".");
            }
            File.Delete("Illuminati.zip");
        }

        private string GetServerVersion()
        {
            string url = APIEndPoint + "/version";
            var webRequest = WebRequest.Create(url);

            // Send the http request and wait for the response
            var responseStream = webRequest.GetResponse().GetResponseStream();

            // Displays the response stream text
            if (responseStream != null)
            {
                using (var streamReader = new StreamReader(responseStream))
                {
                    return streamReader.ReadLine();
                }
            }
            return "";
        }

        private string GetLocalVersion()
        {
            try
            {
                using (StreamReader sr = new StreamReader(".version"))
                {
                    return sr.ReadToEnd();
                }
            }
            catch (Exception e)
            {
                Console.WriteLine(e.Message);
            }
            return "";
        }

    }
}
