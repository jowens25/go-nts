package main

import "log"

func read_ClkClockCoreType_values(core Core) {

	var clock = Clock{

		ControlReg:               0x00000000,
		StatusReg:                0x00000004,
		SelectReg:                0x00000008,
		VersionReg:               0x0000000C,
		TimeValueLReg:            0x00000010,
		TimeValueHReg:            0x00000014,
		TimeAdjValueLReg:         0x00000020,
		TimeAdjValueHReg:         0x00000024,
		OffsetAdjValueReg:        0x00000030,
		OffsetAdjIntervalReg:     0x00000034,
		DriftAdjValueReg:         0x00000040,
		DriftAdjIntervalReg:      0x00000044,
		InSyncThresholdReg:       0x00000050,
		ServoOffsetFactorPReg:    0x00000060,
		ServoOffsetFactorIReg:    0x00000064,
		ServoDriftFactorPReg:     0x00000068,
		ServoDriftFactorIReg:     0x0000006C,
		StatusOffsetReg:          0x00000070,
		StatusDriftReg:           0x00000074,
		StatusOffsetFractionsReg: 0x00000078,
		StatusDriftFractionsReg:  0x0000007C,
	}

	var temp_data int64 = 0x00000000
	//var test int64 = 0x40000001
	//write_reg(core.BaseAddrLReg+clock.ControlReg, &test)
	// enabled
	if read_reg(core.BaseAddrLReg+clock.ControlReg, &temp_data) == 0 {
		//log.Println(temp_data)
		if (temp_data & 0x00000001) == 0 {
			log.Println("ClkClock Enabled: False")
		} else {
			log.Println("ClkClock Enabled: True")
		}
	} else {
		log.Println("ClkClock Enabled: False")
	}
	temp_data = 0x40000000

	// in sync
	if read_reg(core.BaseAddrLReg+clock.StatusReg, &temp_data) == 0 {
		if (temp_data & 0x00000001) == 0 {
			log.Println("ClkClockInSyncValue: False")

		} else {
			log.Println("ClkClockInSyncValue: True")
		}

		if (temp_data & 0x00000002) == 0 {
			log.Println("ClkClockInHoldoverValue: False")

		} else {
			log.Println("ClkClockInHoldoverValue: False")
		}
	} else {
		log.Println("ClkClockInSyncValue: NA")
		log.Println("ClkClockInHoldoverValue: NA")
	}

	// in sync Threshold
	if read_reg(core.BaseAddrLReg+clock.InSyncThresholdReg, &temp_data) == 0 {
		log.Println("ClkClockInSyncThresholdValue: ", temp_data)
	} else {
		log.Println("ClkClockInSyncThresholdValue: NA")
	}

	// offset
	if read_reg(core.BaseAddrLReg+clock.OffsetAdjValueReg, &temp_data) == 0 {

		temp_offset := temp_data & 0x7FFFFFFF
		if (temp_data & 0x80000000) != 0 {
			temp_offset = -1 * temp_offset
		}
		log.Println("ClkClockOffsetValue: ", temp_offset)

		if read_reg(core.BaseAddrLReg+clock.OffsetAdjIntervalReg, &temp_data) == 0 {
			log.Println("ClkClockOffsetIntervalValue: ", temp_data)
		} else {
			log.Println("ClkClockOffsetValue: NA")

			log.Println("ClkClockOffsetIntervalValue: NA")

			log.Println("ClkClockOffsetAdjCheckBox: False")
		}
	} else {
		log.Println("ClkClockOffsetValue: NA")

		log.Println("ClkClockOffsetIntervalValue: NA")

		log.Println("ClkClockOffsetAdjCheckBox: False")
	}

	/*


	   	temp_data = 0x40000000;
	       if(true == ui->ClkClockEnableCheckBox->isChecked())
	       {
	           temp_data |= 0x00000001;
	       }
	       if (0 ==  ucm->com_lib.write_reg(temp_addr + Ucm_ClkClock_ControlReg, temp_data))
	       {
	           for (int i = 0; i < 10; i++)
	           {
	               if(0 ==  ucm->com_lib.read_reg(temp_addr + Ucm_ClkClock_ControlReg, temp_data))
	               {
	                   if ((temp_data & 0x80000000) != 0)
	                   {
	                       // seconds
	                       if (0 ==  ucm->com_lib.read_reg(temp_addr + Ucm_ClkClock_TimeValueHReg, temp_data))
	                       {
	                           ui->ClkClockSecondsValue->setText(QString::number(temp_data));
	                       }
	                       else
	                       {
	                           ui->ClkClockSecondsValue->setText("NA");
	                       }

	                       // date
	                       QDateTime temp_date;
	                       temp_date.setSecsSinceEpoch(temp_data);
	                       ui->ClkClockDateValue->setText(temp_date.toUTC().toString("dd.MM.yyyy hh:mm:ss"));


	                       // nanoseconds
	                       if (0 ==  ucm->com_lib.read_reg(temp_addr + Ucm_ClkClock_TimeValueLReg, temp_data))
	                       {
	                           ui->ClkClockNanosecondsValue->setText(QString::number(temp_data));
	                       }
	                       else
	                       {
	                           ui->ClkClockNanosecondsValue->setText("NA");
	                       }

	                       break;
	                   }
	                   else if (i == 9)
	                   {
	                       cout << "ERROR: " << "read did not complete" << endl;
	                       ui->ClkClockSecondsValue->setText("NA");
	                       ui->ClkClockNanosecondsValue->setText("NA");
	                   }

	               }
	               else
	               {
	                   ui->ClkClockSecondsValue->setText("NA");
	                   ui->ClkClockNanosecondsValue->setText("NA");
	               }
	           }
	       }
	       else
	       {
	           ui->ClkClockSecondsValue->setText("NA");
	           ui->ClkClockNanosecondsValue->setText("NA");
	       } */

}
