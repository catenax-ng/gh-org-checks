import {Component, OnInit} from '@angular/core';
import {GetService} from "./services/get.service";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'GitHub Check Results';

  testSuccess: boolean;
  allData: any;
  testTimeStamp: string;
  repoTests: Array<any>;
  repoReports: Array<any>;

  constructor(private getService: GetService) {
    this.testSuccess = false;
    this.testTimeStamp = "Not available";
    this.repoTests = [];
    this.repoReports = [];
  }

  ngOnInit(): void {
    this.getService.getData().subscribe(
      data => {
        this.allData = data;
        this.testSuccess = true;
        this.testTimeStamp = this.allData.LastTestTime;
        this.repoTests = this.allData.RepositoriesReports[0].RepositoryReport;
        this.repoReports = this.allData.RepositoriesReports
      },
      error => {console.error('There was an error!', error), this.testSuccess = false}
    )



  }


}
